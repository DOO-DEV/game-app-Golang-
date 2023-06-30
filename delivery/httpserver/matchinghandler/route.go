package matchinghandler

import (
	"game-app/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	m := e.Group("/matching")

	m.POST("/add-to-waiting-list", h.AddToWaitingList, middleware.Auth(h.authSvc, h.authConfig))
}
