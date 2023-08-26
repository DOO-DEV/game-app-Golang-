package questionvalidator

import (
	"fmt"
	"game-app/param"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (v Validator) ValidateGetQuestionsByCategoryRequest(req param.GetQuestionsByCategoryRequest) (map[string]string, error) {
	const op = "questionvalidator.ValidateGetQuestionsByCategoryRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.CategoryID, validation.Required, is.Int, validation.By(v.doesCategoryExist)),
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

func (v Validator) doesCategoryExist(value interface{}) error {
	id := value.(uint)

	if _, err := v.repo.GetCategoryByID(id); err != nil {
		return fmt.Errorf(errmsg.ErrorMsgNotFound)
	}

	return nil
}
