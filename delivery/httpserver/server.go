package httpserver

import (
	"fmt"
	"game-app/config"
	"game-app/delivery/httpserver/backoffice_user_handler"
	"game-app/delivery/httpserver/game_websocket"
	"game-app/delivery/httpserver/matchinghandler"
	"game-app/delivery/httpserver/userhandler"
	"game-app/logger"
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backoffice_user_service"
	"game-app/service/matchingservice"
	"game-app/service/presenceservice"
	"game-app/service/userservice"
	"game-app/service/validator/matchingvalidator"
	"game-app/service/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	Router                *echo.Echo
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backoffice_user_handler.Handler
	matchingHandler       matchinghandler.Handler
	gameWsHandler         game_websocket.Handler
}

func New(config config.Config,
	authSvc authservice.Service,
	userSvc userservice.Service,
	userValidator uservalidator.Validator,
	backofficeUserSvc backoffice_user_service.Service,
	authorizationSvc authorizationservice.Service,
	matchingSvc matchingservice.Service,
	matchingValidator matchingvalidator.Validator,
	presenceSvc presenceservice.Service,
) Server {
	return Server{
		Router:                echo.New(),
		config:                config,
		userHandler:           userhandler.New(config.Auth, authSvc, userSvc, userValidator, presenceSvc),
		backofficeUserHandler: backoffice_user_handler.New(config.Auth, authSvc, backofficeUserSvc, authorizationSvc),
		matchingHandler:       matchinghandler.New(config.Auth, authSvc, matchingSvc, matchingValidator, presenceSvc),
		gameWsHandler:         game_websocket.New(config.Auth, authSvc, config.Redis),
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogError:         true,
		LogHost:          true,
		LogMethod:        true,
		LogContentLength: true,
		LogLatency:       true,
		LogResponseSize:  true,
		LogRequestID:     true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Logger.Named("http-server").Info("request",
				zap.String("host", v.Host),
				zap.String("content-length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("methode", v.Method),
				zap.Duration("latency", v.Latency),
				zap.Any("error", v.Error),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response-size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.String("request_id", v.RequestID),
			)

			return nil
		},
	}))

	s.Router.GET("/health-check", s.healthCheck)

	s.userHandler.SetRoutes(s.Router)
	s.backofficeUserHandler.SetRoutes(s.Router)
	s.matchingHandler.SetRoutes(s.Router)
	s.gameWsHandler.SetRoutes(s.Router)

	s.Router.Logger.Fatal(s.Router.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}
