package app

import (
	"context"
	"fmt"
	presenceClient "game-app/adapter/presence"
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/mysql"
	"game-app/repository/mysql/access_control"
	mysqlquestion "game-app/repository/mysql/question"
	mysqluser "game-app/repository/mysql/user"
	"game-app/repository/redis/matchign"
	"game-app/repository/redis/redispresence"
	"game-app/scheduler"
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backoffice_user_service"
	"game-app/service/matchingservice"
	"game-app/service/presenceservice"
	"game-app/service/questionservice"
	"game-app/service/userservice"
	"game-app/service/validator/matchingvalidator"
	"game-app/service/validator/questionvalidator"
	"game-app/service/validator/uservalidator"
	"github.com/spf13/cobra"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"time"
)

type setupSvc struct {
	authSvc           authservice.Service
	userSvc           userservice.Service
	userValidator     uservalidator.Validator
	backofficeUserSvc backoffice_user_service.Service
	authorizationSvc  authorizationservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
	presenceSvc       presenceservice.Service
	questionSvc       questionservice.Service
	questionValidator questionvalidator.Validator
}

func main(cfg config.Config) {
	// curl http://localhost:8099/debug/pprof/gorouitine --output goroutine.tar
	// go tool pprof -http=:8099 ./goroutine.tar
	go http.ListenAndServe(":8099", nil)

	// TODO - add struct and add these returned items as struct fields
	services := setupServices(cfg)

	done := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		sch := scheduler.New(services.matchingSvc, cfg.Scheduler)
		sch.Start(done, &wg)
	}()

	server := httpserver.New(cfg, services.authSvc, services.userSvc, services.userValidator,
		services.backofficeUserSvc, services.authorizationSvc,
		services.matchingSvc, services.matchingValidator,
		services.presenceSvc, services.questionValidator, services.questionSvc)
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

func setupServices(cfg config.Config) setupSvc {
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

	presenceRpo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(presenceRpo, cfg.PresenceService)

	presenceAdaptor := presenceClient.New(":8086")
	matchingSvc := matchingservice.New(matchingRepo, cfg.MatchingService, presenceAdaptor, redisAdapter)

	matchingValidator := matchingvalidator.New()

	questionRepo := mysqlquestion.New(MysqlRepo)
	questionValidator := questionvalidator.New(questionRepo)
	questionSvc := questionservice.New(questionRepo)

	return setupSvc{
		authSvc:           authSvc,
		userSvc:           userSvc,
		userValidator:     userValidator,
		backofficeUserSvc: backofficeUserSvc,
		authorizationSvc:  authorizationSvc,
		matchingSvc:       matchingSvc,
		matchingValidator: matchingValidator,
		presenceSvc:       presenceSvc,
		questionSvc:       questionSvc,
		questionValidator: questionValidator,
	}
}

func New(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "app",
		Short: "run core app",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}
}
