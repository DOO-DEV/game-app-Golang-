package backoffice_user_handler

import (
	"game-app/delivery/httpserver/middleware"
	"game-app/entity"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("backoffice/users")

	userGroup.GET("", h.ListUsers, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))
}
