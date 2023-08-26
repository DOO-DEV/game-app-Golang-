package qusetion

import (
	"game-app/delivery/httpserver/middleware"
	"game-app/entity"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	m := e.Group("/question")

	m.POST("", h.CreateNewQuestion, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))
}
