package config

import (
	"game-app/repository/mysql"
	"game-app/service/authservice"
)

type Config struct {
	Auth       authservice.Config
	HTTPServer HTTPServer
	MySql      mysql.Config
}

type HTTPServer struct {
	Port int
}
