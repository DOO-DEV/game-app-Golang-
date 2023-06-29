package main

import (
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/repository/mysql/access_control"
	mysqluser "game-app/repository/mysql/user"
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backoffice_user_service"
	"game-app/service/userservice"
	"game-app/service/validator/uservalidator"
)

func main() {

	cfg := config.New()

	// TODO - add command for migrations
	mgr := migrator.New(cfg.MySql)
	mgr.Up()

	authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc)

	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service,
	uservalidator.Validator, backoffice_user_service.Service, authorizationservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	MysqlRepo := mysql.New(cfg.MySql)

	userMySql := mysqluser.New(MysqlRepo)
	userSvc := userservice.New(authSvc, userMySql)
	userValidator := uservalidator.New(userMySql)

	backofficeUserSvc := backoffice_user_service.New()

	aclMySql := mysqlaccesscontrol.New(MysqlRepo)
	authorizationSvc := authorizationservice.New(aclMySql)

	return authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc
}
