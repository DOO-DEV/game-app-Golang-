package backoffice_user_handler

import (
	"game-app/pkg/errmsg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) ListUsers(c echo.Context) error {
	list, err := h.backofficeUserSvc.ListAllUsers()
	if err != nil {
		msg, code := httpmsg.Error(err)

		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": list,
	})
}
