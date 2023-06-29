package middleware

import (
	"fmt"
	"game-app/entity"
	"game-app/pkg/claims"
	"game-app/pkg/errmsg"
	"game-app/service/authorizationservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AccessCheck(service authorizationservice.Service, permissions ...entity.PermissionTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := claims.GetClaimsFromEchoContext(c)
			isAllowd, err := service.CheckAccess(claims.UserID, claims.Role, permissions...)
			fmt.Println("error", err)
			if err != nil {
				// TODO - log unexpected error
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errmsg.ErrorMsgSomethingWentWrong,
				})
			}
			if !isAllowd {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": errmsg.ErrorMsgAccessDenied,
				})
			}
			return next(c)
		}
	}
}
