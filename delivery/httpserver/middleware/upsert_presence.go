package middleware

import (
	"game-app/param"
	"game-app/pkg/claims"
	"game-app/pkg/errmsg"
	"game-app/pkg/timestamp"
	"game-app/service/presenceservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpsertPresence(service presenceservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := claims.GetClaimsFromEchoContext(c)

			_, err := service.UpsertPresence(c.Request().Context(), param.UpsertPresenceRequest{
				UserID:    claims.UserID,
				Timestamp: timestamp.Now(),
			})

			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errmsg.ErrorMsgSomethingWentWrong,
				})
			}
			return next(c)
		}
	}
}
