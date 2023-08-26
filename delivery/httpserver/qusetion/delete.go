package qusetion

import (
	"game-app/param"
	"game-app/pkg/errmsg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) DeleteQuestion(c echo.Context) error {
	var req param.DeleteQuestionRequest

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	if fieldErrors, err := h.questionValidator.ValidateDeleteQuestionRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"errors":  fieldErrors,
			"message": msg,
		})
	}

	res, err := h.questionSvc.DeleteQuestion(req)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, res)
}
