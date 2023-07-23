package game_websocket

import (
	"game-app/adapter/redis"
	"game-app/service/authservice"
)

type Handler struct {
	authSvc     authservice.Service
	authConfig  authservice.Config
	redisconfig redis.Config
}

func New(authConfig authservice.Config, authSvc authservice.Service, redisConfig redis.Config) Handler {
	return Handler{
		authSvc:     authSvc,
		authConfig:  authConfig,
		redisconfig: redisConfig,
	}
}
