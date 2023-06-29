package backoffice_user_handler

import (
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backoffice_user_service"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	backofficeUserSvc backoffice_user_service.Service
	authorizationSvc  authorizationservice.Service
}

func New(authConfig authservice.Config, authSvc authservice.Service,
	backofficeUserSvc backoffice_user_service.Service, authorizationSvc authorizationservice.Service,
) Handler {
	return Handler{
		authConfig:        authConfig,
		authSvc:           authSvc,
		backofficeUserSvc: backofficeUserSvc,
		authorizationSvc:  authorizationSvc,
	}
}
