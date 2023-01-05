package cloud_computer

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"strings"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		DB:         0,
		MaxRetries: 5,
	})

	if status := client.Ping(context.TODO()); status.Err() != nil {
		log.Fatal(status.Err())
	}

	return client
}

// TODO: rename Read to Subscribe
func ReadAsyncRedis(ctx context.Context, client *redis.Client, name string) (status <-chan bool) {
	sChannel := make(chan bool, 1)

	//log.Println("running read async redis", name)
	sub := client.Subscribe(ctx, name)
	if err := sub.Ping(ctx); err != nil {
		log.Println(name + " die due to error on subscribe")
		panic(err)
	}
	go func(sub *redis.PubSub) {
		defer sub.Close()
		defer sub.Unsubscribe(ctx, name)

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

func addChildren(ctx context.Context, client *redis.Client, name string) {
	parents := strings.Split(name, ".")
	client.SAdd(ctx, parents[0]+".children", strings.Join(parents[:len(parents)-1], "."))
}

func addInput(ctx context.Context, client *redis.Client, gateName, name string) {
	//client.SAdd(gateName+".inputs", name)

	parents := strings.Split(gateName, ".")
	client.SAdd(ctx, strings.Join(parents[:len(parents)-1], ".")+".inputs", name)
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

			err := writeRedis(ctx, client, name+".status", s)
			if err != nil {
				panic(err)
			}

			//fmt.Printf("message at channel %s = %v\n", name, s)

			intCmd := client.Publish(ctx, name, data)
			if intCmd.Err() != nil {
				panic(intCmd.Err())
			}
		}
	}(client, name)

	return sChannel
}

func addOutput(ctx context.Context, client *redis.Client, gateName, name string) {
	//client.SAdd(gateName+".outputs", name)

	parents := strings.Split(gateName, ".")
	client.SAdd(ctx, strings.Join(parents[:len(parents)-1], ".")+".outputs", name)
	//client.SAdd(parents[0]+".outputs", name)
}

func ReadRedis(ctx context.Context, client *redis.Client, name string) (value bool, err error) {
	v, err := client.Get(ctx, name).Result()
	if err != nil && err != redis.Nil {
		return
	}

	if v == "1" {
		return true, nil
	}
	return false, nil
}

func writeRedis(ctx context.Context, client *redis.Client, name string, value bool) (err error) {
	return client.Set(ctx, name, value, 0).Err()
}

func deleteRedis(ctx context.Context, client *redis.Client, name string) {
	err := client.Del(ctx, name).Err()
	if err != nil {
		panic(err)
	}
}
