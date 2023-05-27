package main

import (
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/service/authservice"
	"game-app/service/userservice"
	"game-app/service/validator/uservalidator"
)

func main() {

	cfg := config.New()

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
