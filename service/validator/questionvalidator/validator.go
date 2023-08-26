package questionvalidator

import "game-app/entity"

type Repository interface {
	GetQuestionByID(id uint) (entity.Question, error)
	GetCategoryByID(id uint) (entity.Category, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
