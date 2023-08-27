package questionservice

import (
	"fmt"
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) DeleteQuestion(req param.DeleteQuestionRequest) (param.DeleteQuestionResponse, error) {
	const op = "questionservice.DeleteQuestion"

	if err := s.repo.DeleteQuestion(req.ID); err != nil {
		return param.DeleteQuestionResponse{}, richerror.New(op).WithErr(err)
	}

	return param.DeleteQuestionResponse{
		Message: fmt.Sprintf("question with id %d successfully deleted", req.ID),
	}, nil
}
