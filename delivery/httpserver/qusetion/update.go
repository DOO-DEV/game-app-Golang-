package qusetion

import (
	"game-app/param"
	"game-app/pkg/errmsg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h Handler) UpdateQuestion(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	var req param.UpdateQuestionRequest
	req.ID = uint(id)

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	if fieldErrors, err := h.questionValidator.ValidateUpdateQuestionRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"errors":  fieldErrors,
			"message": msg,
		})
	}

	res, err := h.questionSvc.UpdateQuestion(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
