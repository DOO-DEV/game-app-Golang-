package question

import (
	"game-app/param"
	"game-app/pkg/errmsg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h Handler) GetByCategory(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	var req param.GetQuestionsByCategoryRequest
	req.CategoryID = uint(id)

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	if fieldErrors, err := h.questionValidator.ValidateGetQuestionsByCategoryRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"errors":  fieldErrors,
			"message": msg,
		})
	}

	res, err := h.questionSvc.GetQuestionsByCategory(req)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, res)
}
