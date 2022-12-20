package cloud_computer

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func ReadAsyncRedis(ctx context.Context, client *redis.Client, name string) (status <-chan bool) {
	sChannel := make(chan bool, 1)

	log.Println("running read async redis")
	sub := client.Subscribe(name)
	go func(sub *redis.PubSub) {
		defer sub.Close()
		defer sub.Unsubscribe(name)

		for msg := range sub.Channel() {
			fmt.Printf("message at channel %s = %s\n", msg.Channel, msg.Payload)

			if msg.Payload == "1" {
				sChannel <- true
			} else {
				sChannel <- false
			}
		}
	}(sub)

	return sChannel
}

func WriteAsyncRedis(ctx context.Context, client *redis.Client, name string) (status chan<- bool) {
	sChannel := make(chan bool, 1)

	go func(client *redis.Client, name string) {
		for s := range sChannel {
			data := 0
			if s {
				data = 1
			}

			intCmd := client.Publish(name, data)
			if intCmd.Err() != nil {
				panic(intCmd.Err())
			}

			err := writeRedis(ctx, client, name+".status", s)
			if err != nil {
				panic(err)
			}
		}
	}(client, name)

	return sChannel
}

func writeRedis(ctx context.Context, client *redis.Client, name string, value bool) (err error) {
	return client.Set(name, value, 0).Err()
}

func deleteRedis(ctx context.Context, client *redis.Client, name string) {
	err := client.Del(name).Err()
	if err != nil {
		panic(err)
	}
}
