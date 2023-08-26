package questionservice

import (
	"game-app/entity"
	"game-app/param"
)

type Repository interface {
	GetQuestionByID(id uint) (entity.Question, error)
	InsertQuestion(question entity.Question) (entity.Question, error)
	UpdateQuestion(question entity.Question) (entity.Question, error)
	DeleteQuestion(id uint) error
	GetQuestionsByID(category entity.Category) ([]entity.Question, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetQuestion(req param.GetQuestionRequest) (param.GetQuestionResponse, error) {
	return param.GetQuestionResponse{}, nil
}

func (s Service) CreateNewQuestion(req param.CreateNewQuestionRequest) (param.CreateNewGameResponse, error) {

	return param.CreateNewGameResponse{}, nil
}

func (s Service) UpdateQuestion(req param.UpdateQuestionRequest) (param.UpdateQuestionResponse, error) {
	return param.UpdateQuestionResponse{}, nil
}

func (s Service) DeleteQuestion(req param.DeleteQuestionRequest) (param.DeleteQuestionResponse, error) {
	return param.DeleteQuestionResponse{}, nil
}

func (s Service) GetQuestionsByCategory(request param.GetQuestionsByCategoryRequest) (param.GetQuestionsByCategoryResponse, error) {
	return param.GetQuestionsByCategoryResponse{}, nil
}
