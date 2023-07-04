package main

import (
	"context"
	"fmt"
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/repository/mysql/access_control"
	mysqluser "game-app/repository/mysql/user"
	"game-app/repository/redis/matchign"
	"game-app/repository/redis/redispresence"
	"game-app/scheduler"
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backoffice_user_service"
	"game-app/service/matchingservice"
	"game-app/service/presenceservice"
	"game-app/service/userservice"
	"game-app/service/validator/matchingvalidator"
	"game-app/service/validator/uservalidator"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	cfg := config.New()

	// TODO - add command for migrations
	mgr := migrator.New(cfg.MySql)
	mgr.Up()

	// TODO - add struct and add these returned items as struct fields
	authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingValidator, presenceSvc := setupServices(cfg)

	done := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		sch := scheduler.New(matchingSvc, cfg.Scheduler)
		sch.Start(done, &wg)
	}()

	server := httpserver.New(cfg, authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc,
		matchingSvc, matchingValidator, presenceSvc)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.Application.GracefullShutDownTimeout)
	defer cancel()

	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("received interrupt signal, shutting down gracefully...")
	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("error while shutting down echo http server: ", err)
	}
	done <- true
	time.Sleep(cfg.Application.GracefullShutDownTimeout)

	// TODO - the context doesn't wait for scheduler to finish its job..
	<-ctxWithTimeout.Done()

	wg.Wait()
}

func setupServices(cfg config.Config) (
	authservice.Service,
	userservice.Service,
	uservalidator.Validator,
	backoffice_user_service.Service,
	authorizationservice.Service,
	matchingservice.Service,
	matchingvalidator.Validator,
	presenceservice.Service,
) {
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

	presenceRpo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(presenceRpo, cfg.PresenceService)

	return authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingValidator, presenceSvc
}
