package questionvalidator

import (
	"fmt"
	"game-app/param"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateDeleteQuestionRequest(req param.DeleteQuestionRequest) (map[string]string, error) {
	const op = "questionvalidator.ValidateDeleteQuestionRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.By(v.doesQuestionExist)),
	); err != nil {
		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgInvalidInput).WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req})
	}

	return nil, nil
}

func (v Validator) doesQuestionExist(value interface{}) error {
	id := value.(uint)

	_, err := v.repo.GetQuestionByID(id)
	if err != nil {
		return fmt.Errorf(errmsg.ErrorMsgNotFound)
	}

	return nil
}
