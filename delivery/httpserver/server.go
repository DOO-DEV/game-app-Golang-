package httpserver

import (
	"fmt"
	"game-app/config"
	"game-app/delivery/httpserver/backoffice_user_handler"
	"game-app/delivery/httpserver/userhandler"
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backoffice_user_service"
	"game-app/service/userservice"
	"game-app/service/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backoffice_user_handler.Handler
}

func New(config config.Config,
	authSvc authservice.Service,
	userSvc userservice.Service,
	userValidator uservalidator.Validator,
	backofficeUserSvc backoffice_user_service.Service,
	authorizationSvc authorizationservice.Service,
) Server {
	return Server{
		config:                config,
		userHandler:           userhandler.New(config.Auth, authSvc, userSvc, userValidator),
		backofficeUserHandler: backoffice_user_handler.New(config.Auth, authSvc, backofficeUserSvc, authorizationSvc),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)

	s.userHandler.SetRoutes(e)
	s.backofficeUserHandler.SetRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
