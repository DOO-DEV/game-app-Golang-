package main

import (
	"context"
	"fmt"
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/pkg/protobufencoder"
)

func main() {
	cfg := config.New()

	redisAdapter := redis.New(cfg.Redis)

	topic := "matching.users_matched"
	subscriber := redisAdapter.Client().Subscribe(context.Background(), topic)

	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			panic(err)
		}

		payload := protobufencoder.DecodeMatchingUsersMatchedEvent(msg.Payload)

		fmt.Println("Received message from " + msg.Channel + "channel.")
		fmt.Printf("matched users %+v\n", payload)
	}

}
