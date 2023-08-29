package questionservice

import (
	"context"
	"game-app/entity"
	"game-app/param"
)

type Repository interface {
	GetQuestionByID(id uint) (entity.Question, error)
	InsertQuestion(question entity.Question) (entity.Question, error)
	UpdateQuestion(question entity.Question) (entity.Question, error)
	DeleteQuestion(id uint) error
	GetQuestionsByCategory(id uint) ([]entity.Question, error)
}

type AnswerClient interface {
	GetAnswers(ctx context.Context, req param.GetAnswersRequest) (param.GetAnswersResponse, error)
	InsertAnswers(ctx context.Context, req param.InsertAnswersRequest) (param.InsertAnswersResponse, error)
	DeleteAnswer(ctx context.Context, req param.DeleteAnswerRequest) (param.DeleteAnswerResponse, error)
	UpdateAnswer(ctx context.Context, req param.UpdateAnswerRequest) (param.UpdateAnswerResponse, error)
}

type Service struct {
	repo         Repository
	answerClient AnswerClient
}

func New(repo Repository, answerClient AnswerClient) Service {
	return Service{
		repo:         repo,
		answerClient: answerClient,
	}
}
