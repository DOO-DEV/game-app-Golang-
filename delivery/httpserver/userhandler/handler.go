package userhandler

import (
	"game-app/service/authservice"
	"game-app/service/userservice"
	"game-app/service/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
