package main

import (
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/service/authservice"
	"game-app/service/userservice"
	"game-app/service/validator/uservalidator"
	"time"
)

const (
	JwtSignKey            = "jwt_secret"
	AccessSubject         = "at"
	RefreshSubject        = "rt"
	AccessExpirationTime  = time.Hour * 24
	RefreshExpirationTime = time.Hour * 24 * 7
)

func main() {
	cfg := config.Config{
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessExpirationTime,
			RefreshExpirationTime: RefreshExpirationTime,
			AccessSubject:         AccessSubject,
			RefreshSubject:        RefreshSubject,
		},
		HTTPServer: config.HTTPServer{Port: 8080},
		MySql: mysql.Config{
			Username: "gameapp",
			Password: "gameappt0lk2o20",
			Port:     3308,
			Host:     "localhost",
			DBName:   "gameapp_db",
		},
	}

	// TODO - add command for migrations
	mgr := migrator.New(cfg.MySql)
	mgr.Up()

	authSvc, userSvc, userValidator := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator)

	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)
	MysqlRepo := mysql.New(cfg.MySql)
	userSvc := userservice.New(authSvc, MysqlRepo)
	userValidator := uservalidator.New(MysqlRepo)

	return authSvc, userSvc, userValidator
}
