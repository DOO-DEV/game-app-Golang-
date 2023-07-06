package main

import (
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/delivery/grpcserver/presenceserver"
	"game-app/repository/redis/redispresence"
	"game-app/service/presenceservice"
)

func main() {
	// TODO - read config path from command line
	cfg := config.New()

	redisAdapter := redis.New(cfg.Redis)

	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(presenceRepo, cfg.PresenceService)

	server := presenceserver.New(presenceSvc)
	server.Start()
}
