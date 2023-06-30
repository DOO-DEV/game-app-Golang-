package main

import (
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/repository/mysql/access_control"
	mysqluser "game-app/repository/mysql/user"
	"game-app/repository/redis/matchign"
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backoffice_user_service"
	"game-app/service/matchingservice"
	"game-app/service/userservice"
	"game-app/service/validator/matchingvalidator"
	"game-app/service/validator/uservalidator"
)

func main() {

	cfg := config.New()

	// TODO - add command for migrations
	mgr := migrator.New(cfg.MySql)
	mgr.Up()

	// TODO - add struct and add these returned items as struct fields
	authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingValidator := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingValidator)

	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service,
	userservice.Service,
	uservalidator.Validator,
	backoffice_user_service.Service,
	authorizationservice.Service,
	matchingservice.Service,
	matchingvalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)
	MysqlRepo := mysql.New(cfg.MySql)

	userMySql := mysqluser.New(MysqlRepo)
	userSvc := userservice.New(authSvc, userMySql)
	userValidator := uservalidator.New(userMySql)

	backofficeUserSvc := backoffice_user_service.New()
	aclMySql := mysqlaccesscontrol.New(MysqlRepo)
	authorizationSvc := authorizationservice.New(aclMySql)

	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := matchign.New(redisAdapter)
	matchingSvc := matchingservice.New(matchingRepo, cfg.MatchingService)
	matchingValidator := matchingvalidator.New()

	return authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingValidator
}
