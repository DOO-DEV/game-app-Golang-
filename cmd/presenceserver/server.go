package main

import (
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/delivery/grpcserver/presenceserver"
	"game-app/repository/redis/redispresence"
	"game-app/service/presenceservice"
)

func main() {
	cfg := config.New()
	redisAdapter := redis.New(cfg.Redis)
	presenceRpo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(presenceRpo, cfg.PresenceService)

	server := presenceserver.New(presenceSvc)
	server.Start()
}
