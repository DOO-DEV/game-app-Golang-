package game_websocket

import (
	"game-app/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/game/ws")

	g.GET("", h.GameWs, middleware.Auth(h.authSvc, h.authConfig))
}
