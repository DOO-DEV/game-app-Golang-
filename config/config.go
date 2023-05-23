package config

import (
	"game-app/repository/mysql"
	"game-app/service/authservice"
)

type Config struct {
	Auth       authservice.Config `koanf:"auth"`
	HTTPServer HTTPServer         `koanf:"http_server"`
	MySql      mysql.Config       `koanf:"mysql"`
	Debug      bool               `koanf:"debug"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}
