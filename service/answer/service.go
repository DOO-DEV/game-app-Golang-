package answer

import (
	"fmt"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
)

type Repository interface {
	InsertAnswers(answers []entity.PossibleAnswer) error
	UpdateAnswer(answer entity.PossibleAnswer) (entity.PossibleAnswer, error)
	DeleteAnswer(id uint) error
}
type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) InsertAnswers(req param.InsertAnswersRequest) (param.InsertAnswersResponse, error) {
	const op = "answerservice.InsertAnswers"

	answers := make([]entity.PossibleAnswer, 0)
	for _, ans := range req.Data {
		answer := entity.PossibleAnswer{
			QuestionID: req.QuestionID,
			Text:       ans.Text,
			Choice:     entity.PossibleAnswerChoice(ans.Choice),
		}
		answers = append(answers, answer)
	}

	if err := s.repo.InsertAnswers(answers); err != nil {
		return param.InsertAnswersResponse{}, richerror.New(op).WithErr(err)
	}

	return param.InsertAnswersResponse{Message: "answers successfully added"}, nil
}

func (s Service) UpdateAnswer(req param.UpdateAnswerRequest) (param.UpdateAnswerResponse, error) {
	const op = "answerservice.UpdateAnswer"

	newAnswer := entity.PossibleAnswer{
		ID:         req.ID,
		QuestionID: req.QuestionID,
		Text:       req.Data.Text,
		Choice:     entity.PossibleAnswerChoice(req.Data.Choice),
	}

	answer, err := s.repo.UpdateAnswer(newAnswer)
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

func (s Service) DeleteAnswer(req param.DeleteAnswerRequest) (param.DeleteAnswerResponse, error) {
	const op = "answerservice.DeleteAnswer"

	if err := s.repo.DeleteAnswer(req.ID); err != nil {
		return param.DeleteAnswerResponse{}, richerror.New(op).WithErr(err)
	}

	return param.DeleteAnswerResponse{Message: fmt.Sprintf("answer with id %d deleted", req.ID)}, nil
}
