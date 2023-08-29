package answerservice

import (
	"context"
	"fmt"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
)

type Repository interface {
	InsertAnswers(ctx context.Context, answers []entity.PossibleAnswer) error
	UpdateAnswer(ctx context.Context, answer entity.PossibleAnswer) (entity.PossibleAnswer, error)
	DeleteAnswer(ctx context.Context, id uint) error
	GetAnswers(ctx context.Context, id uint) ([]entity.PossibleAnswer, error)
}
type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) GetAnswers(ctx context.Context, req param.GetAnswersRequest) (param.GetAnswersResponse, error) {
	const op = "answer.GetAnswers"

	answers, err := s.repo.GetAnswers(ctx, req.QuestionID)
	if err != nil {
		return param.GetAnswersResponse{}, richerror.New(op).WithErr(err)
	}

	res := make([]param.Answer, 0)
	for _, item := range answers {
		res = append(res, param.Answer{
			Text:   item.Text,
			Choice: uint(item.Choice),
		})
	}

	return param.GetAnswersResponse{
		QuestionID: req.QuestionID,
		Data:       res,
	}, nil
}

func (s Service) InsertAnswers(ctx context.Context, req param.InsertAnswersRequest) (param.InsertAnswersResponse, error) {
	const op = "answer.InsertAnswers"

	answers := make([]entity.PossibleAnswer, 0)
	for _, ans := range req.Data {
		answer := entity.PossibleAnswer{
			QuestionID: req.QuestionID,
			Text:       ans.Text,
			Choice:     entity.PossibleAnswerChoice(ans.Choice),
		}
		answers = append(answers, answer)
	}

	if err := s.repo.InsertAnswers(ctx, answers); err != nil {
		return param.InsertAnswersResponse{}, richerror.New(op).WithErr(err)
	}

	return param.InsertAnswersResponse{}, nil
}

func (s Service) UpdateAnswer(ctx context.Context, req param.UpdateAnswerRequest) (param.UpdateAnswerResponse, error) {
	const op = "answer.UpdateAnswer"

	newAnswer := entity.PossibleAnswer{
		ID:         req.ID,
		QuestionID: req.QuestionID,
		Text:       req.Data.Text,
		Choice:     entity.PossibleAnswerChoice(req.Data.Choice),
	}

	answer, err := s.repo.UpdateAnswer(ctx, newAnswer)
	if err != nil {
		return param.UpdateAnswerResponse{}, richerror.New(op).WithErr(err)
	}

	return param.UpdateAnswerResponse{
		ID: answer.ID,
		Data: param.Answer{
			Text:   answer.Text,
			Choice: uint(answer.Choice),
		},
	}, nil
}

func (s Service) DeleteAnswer(ctx context.Context, req param.DeleteAnswerRequest) (param.DeleteAnswerResponse, error) {
	const op = "answer.DeleteAnswer"

	if err := s.repo.DeleteAnswer(ctx, req.ID); err != nil {
		return param.DeleteAnswerResponse{}, richerror.New(op).WithErr(err)
	}

	return param.DeleteAnswerResponse{Message: fmt.Sprintf("answer with id %d deleted", req.ID)}, nil
}
