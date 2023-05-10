package httpmsg

import (
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"net/http"
)

func Error(err error) (string, int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.Message()
		code := kindToHTTPStatusCode(re.Kind())
		if code >= 500 {
			msg = errmsg.ErrorMsgSomethingWentWrong
		}
		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func kindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
