package questionservice

import (
	"context"
	"fmt"
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) DeleteQuestion(_ context.Context, req param.DeleteQuestionRequest) (param.DeleteQuestionResponse, error) {
	const op = "questionservice.DeleteQuestion"

	if err := s.repo.DeleteQuestion(req.ID); err != nil {
		return param.DeleteQuestionResponse{}, richerror.New(op).WithErr(err)
	}

	// ON DELETE CASCADE on database set ==> it will be delete all child (has foregien key) rows in child table;

	return param.DeleteQuestionResponse{
		Message: fmt.Sprintf("question with id %d successfully deleted", req.ID),
	}, nil
}
