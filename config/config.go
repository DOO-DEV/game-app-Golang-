package config

import (
	"game-app/adapter/redis"
	"game-app/repository/mysql"
	"game-app/service/authservice"
	"game-app/service/matchingservice"
)

type Config struct {
	Auth            authservice.Config     `koanf:"auth"`
	HTTPServer      HTTPServer             `koanf:"http_server"`
	MySql           mysql.Config           `koanf:"mysql"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
	Redis           redis.Config           `koanf:"redis"`
	Debug           bool                   `koanf:"debug"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}
