package cloud_computer

import (
	"bytes"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func readAsyncKafka(ctx context.Context, topic string, partition int) (status chan bool) {
	status = make(chan bool, 1)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	go func() {
		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				break
			}
			fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

			if bytes.Equal(m.Value, trueBytes) {
				status <- true
			} else {
				status <- false
			}
		}

		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}()

	return
}

var falseBytes = []byte("false")
var trueBytes = []byte("true")

func writeAsyncKafka(ctx context.Context, topic string, partition int) (status chan bool) {
	status = make(chan bool, 1)

	go func() {
		conn, err := kafka.DialLeader(ctx, "tcp", "localhost:9092", topic, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		for s := range status {
			b := falseBytes
			if s {
				b = trueBytes
			}

			_, err = conn.WriteMessages(
				kafka.Message{Value: []byte(b)},
			)
			if err != nil {
				log.Fatal("failed to write messages:", err)
			}

			if err := conn.Close(); err != nil {
				log.Fatal("failed to close writer:", err)
			}
		}
		//conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	}()
	return
}
