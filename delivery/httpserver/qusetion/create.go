package question

import (
	"game-app/param"
	"game-app/pkg/errmsg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) CreateNewQuestion(c echo.Context) error {
	var req param.CreateNewQuestionRequest

	if err := c.Bind(&req.Data); err != nil {
		return echo.ErrBadRequest
	}
	if fieldErrors, err := h.questionValidator.ValidateCreateNewQuestionRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"errors":  fieldErrors,
			"message": msg,
		})
	}

	res, err := h.questionSvc.CreateNewQuestion(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
