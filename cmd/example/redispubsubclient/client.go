package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/contract/golang/matching"
	"game-app/entity"
	"game-app/pkg/slice"
	"google.golang.org/protobuf/proto"
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

		payload, err := base64.StdEncoding.DecodeString(msg.Payload)
		if err != nil {
			panic(err)
		}
		pbMu := matching.MatchedUsers{}
		if err := proto.Unmarshal(payload, &pbMu); err != nil {
			panic(err)
		}

		mu := entity.MatchedPlayers{
			Category: entity.Category(pbMu.Category),
			UserIDs:  slice.MapFromUint64ToUint(pbMu.UserIds),
		}

		fmt.Println("Received message from " + msg.Channel + "channel.")
		fmt.Printf("matched users %+v\n", mu)
	}

}
