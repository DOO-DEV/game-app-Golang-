package matchingvalidator

import (
	"fmt"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateAddToWaitingListRequest(req param.AddToWaitingListRequest) (map[string]string, error) {
	const op = "matchingvalidator.AddToWaitingListRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required,
			validation.By(v.isCategoryValid)),
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

func (v Validator) isCategoryValid(value interface{}) error {
	c := value.(entity.Category)

	if !c.IsValid() {
		return fmt.Errorf(errmsg.ErrorMsgCategoryIsNotValid)
	}

	return nil
}
