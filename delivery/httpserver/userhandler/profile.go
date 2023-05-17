package userhandler

import (
	"game-app/pkg/errmsg/httpmsg"
	"game-app/service/userservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	claims, err := h.authSvc.ParseToken(authToken)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	resp, err := h.userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
