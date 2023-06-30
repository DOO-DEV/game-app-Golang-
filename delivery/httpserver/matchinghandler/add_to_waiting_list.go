package matchinghandler

import (
	"game-app/param"
	claim "game-app/pkg/claims"
	"game-app/pkg/errmsg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) AddToWaitingList(c echo.Context) error {
	var req param.AddToWaitingListRequest

	claims := claim.GetClaimsFromEchoContext(c)
	req.UserID = claims.UserID

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if fieldErrors, err := h.matchingValidator.ValidateAddToWaitingListRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"errors":  fieldErrors,
			"message": msg,
		})
	}

	resp, err := h.matchingSvc.AddToWaitingList(req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
