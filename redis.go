package cloud_computer

import (
	"context"
	"github.com/go-redis/redis"
	"strings"
)

func ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

// TODO: rename Read to Subscribe
func ReadAsyncRedis(ctx context.Context, client *redis.Client, name string) (status <-chan bool) {
	sChannel := make(chan bool, 1)

	//log.Println("running read async redis", name)
	sub := client.Subscribe(name)
	go func(sub *redis.PubSub) {
		defer sub.Close()
		defer sub.Unsubscribe(name)

		for msg := range sub.Channel() {
			//fmt.Printf("message at channel %s = %s\n", msg.Channel, msg.Payload)

			if msg.Payload == "1" {
				sChannel <- true
			} else {
				sChannel <- false
			}
		}
	}(sub)

	return sChannel
}

func addInput(client *redis.Client, gateName, name string) {
	//client.SAdd(gateName+".inputs", name)

	parents := strings.Split(gateName, ".")
	client.SAdd(parents[0]+".inputs", name)
}

// TODO: rename Write to Publish
func WriteAsyncRedis(ctx context.Context, client *redis.Client, name string) (status chan<- bool) {
	sChannel := make(chan bool, 1)

	go func(client *redis.Client, name string) {
		for s := range sChannel {
			data := 0
			if s {
				data = 1
			}

			//fmt.Printf("message at channel %s = %v\n", name, s)

			intCmd := client.Publish(name, data)
			if intCmd.Err() != nil {
				panic(intCmd.Err())
			}

			// TODO: output의 값이 status에 덮어씌여지고 있음
			err := writeRedis(ctx, client, name+".status", s)
			if err != nil {
				panic(err)
			}
		}
	}(client, name)

	return sChannel
}

func addOutput(client *redis.Client, gateName, name string) {
	//client.SAdd(gateName+".outputs", name)

	parents := strings.Split(gateName, ".")
	client.SAdd(parents[0]+".outputs", name)
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
