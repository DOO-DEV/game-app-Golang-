package httpserver

import (
	"game-app/dto"
	"game-app/pkg/errmsg/httpmsg"
	"game-app/service/userservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) userRegister(c echo.Context) error {
	// TODO - we should verify by phone number by verification code

	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if fieldErrors, err := s.userValidator.ValidateRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"errors":  fieldErrors,
			"message": msg,
		})
	}

	res, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (s Server) userLogin(c echo.Context) error {
	var req userservice.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	res, err := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (s Server) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	claims, err := s.authSvc.ParseToken(authToken)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	resp, err := s.userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
