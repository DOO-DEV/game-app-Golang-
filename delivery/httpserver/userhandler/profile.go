package userhandler

import (
	"game-app/param"
	"game-app/pkg/claims"
	"game-app/pkg/errmsg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userProfile(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c)

	resp, err := h.userSvc.GetProfile(c.Request().Context(), param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
