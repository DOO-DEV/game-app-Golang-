package question

import (
	"game-app/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	m := e.Group("/question")

	m.POST("", h.CreateNewQuestion, middleware.Auth(h.authSvc, h.authConfig))
	m.DELETE("/:id", h.DeleteQuestion, middleware.Auth(h.authSvc, h.authConfig))
	m.GET("/:id", h.GetQuestion, middleware.Auth(h.authSvc, h.authConfig))
	m.GET("/category/:id", h.GetByCategory, middleware.Auth(h.authSvc, h.authConfig))
}
