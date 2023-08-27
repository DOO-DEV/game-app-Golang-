package questionservice

import (
	"game-app/entity"
)

type Repository interface {
	GetQuestionByID(id uint) (entity.Question, error)
	InsertQuestion(question entity.Question) (entity.Question, error)
	UpdateQuestion(question entity.Question) (entity.Question, error)
	DeleteQuestion(id uint) error
	GetQuestionsByCategory(id uint) ([]entity.Question, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
