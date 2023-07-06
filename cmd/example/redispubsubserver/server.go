package main

import (
	"context"
	"fmt"
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/entity"
	"game-app/pkg/protobufencoder"
)

func main() {
	cfg := config.New()

	redisAdapter := redis.New(cfg.Redis)

	topic := "matching.users_matched"
	mu := entity.MatchedPlayers{
		Category: entity.FootballCategory,
		UserIDs:  []uint{1, 4},
	}

	payloadStr := protobufencoder.EncodeMatchingUsersMatchedEvent(mu)
	if err := redisAdapter.Client().Publish(context.Background(), topic, payloadStr).Err(); err != nil {
		panic(fmt.Sprintf("publish err %v", err))
	}
}
