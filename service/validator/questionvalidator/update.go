package questionvalidator

import (
	"game-app/param"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (v Validator) ValidateUpdateQuestionRequest(req param.UpdateQuestionRequest) (map[string]string, error) {
	const op = "questionvalidator.ValidateUpdateQuestionRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Question, validation.Required, is.Alpha),
		validation.Field(&req.CategoryID, validation.Required, is.Int),
		validation.Field(&req.CorrectAnswerID, validation.Required, is.Int),
		validation.Field(&req.Difficulty, validation.Required, is.Int, validation.Min(1), validation.Max(3)),
		validation.Field(&req.PossibleAnswers, validation.Required),
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
