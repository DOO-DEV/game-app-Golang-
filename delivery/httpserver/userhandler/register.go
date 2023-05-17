package userhandler

import (
	"game-app/param"
	"game-app/pkg/errmsg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userRegister(c echo.Context) error {
	// TODO - we should verify by phone number by verification code

	var req param.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if fieldErrors, err := h.userValidator.ValidateRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"errors":  fieldErrors,
			"message": msg,
		})
	}

	res, err := h.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
