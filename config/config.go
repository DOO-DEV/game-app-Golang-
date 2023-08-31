package config

import (
	"game-app/adapter/redis"
	"game-app/repository/mysql"
	gamescheduler "game-app/scheduler/game"
	"game-app/scheduler/match_waited_user"
	"game-app/service/authservice"
	"game-app/service/gameservice"
	"game-app/service/matchingservice"
	"game-app/service/presenceservice"
	"time"
)

type Config struct {
	Application     Application              `koanf:"application"`
	Auth            authservice.Config       `koanf:"auth"`
	HTTPServer      HTTPServer               `koanf:"http_server"`
	MySql           mysql.Config             `koanf:"mysql"`
	MatchingService matchingservice.Config   `koanf:"matching_service"`
	Redis           redis.Config             `koanf:"redis"`
	PresenceService presenceservice.Config   `koanf:"presence_service"`
	Scheduler       match_waited_user.Config `koanf:"scheduler"`
	Debug           bool                     `koanf:"debug"`
	GameScheduler   gamescheduler.Config     `koanf:"game_scheduler"`
	GameSvc         gameservice.Config       `json:"game_svc"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Application struct {
	GracefullShutDownTimeout time.Duration `koanf:"gracefull_shut_down_timeout"`
}
