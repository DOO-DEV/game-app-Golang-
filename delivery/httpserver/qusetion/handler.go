package question

import (
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backoffice_user_service"
	"game-app/service/questionservice"
	"game-app/service/validator/questionvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	backofficeUserSvc backoffice_user_service.Service
	authorizationSvc  authorizationservice.Service
	questionValidator questionvalidator.Validator
	questionSvc       questionservice.Service
}

func New(
	authConfig authservice.Config,
	authSvc authservice.Service,
	backofficeUserSvc backoffice_user_service.Service,
	authorizationSvc authorizationservice.Service,
	questionValidator questionvalidator.Validator,
	questionSvc questionservice.Service,
) Handler {
	return Handler{
		authConfig:        authConfig,
		authSvc:           authSvc,
		backofficeUserSvc: backofficeUserSvc,
		authorizationSvc:  authorizationSvc,
		questionValidator: questionValidator,
		questionSvc:       questionSvc,
	}
}
