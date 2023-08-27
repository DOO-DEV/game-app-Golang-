package questionservice

import (
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) CreateNewQuestion(req param.CreateNewQuestionRequest) (param.CreateNewQuestionResponse, error) {
	const op = "questionservice.CeateNewQuestion"

	newQuestion := entity.Question{
		Question:        req.Data.Question,
		PossibleAnswers: s.mapToPossibleAnswerEntity(req.Data.PossibleAnswers),
		CorrectAnswerID: req.Data.CorrectAnswerID,
		Difficulty:      entity.QuestionDifficulty(req.Data.Difficulty),
		CategoryID:      req.Data.CategoryID,
	}
	q, err := s.repo.InsertQuestion(newQuestion)
	if err != nil {
		return param.CreateNewQuestionResponse{}, richerror.New(op).WithErr(err)
	}

	res := param.CreateNewQuestionResponse{Data: param.Question{
		ID:              q.ID,
		Question:        q.Question,
		PossibleAnswers: s.mapFromPossibleAnswersEntity(q.PossibleAnswers),
		CorrectAnswerID: q.CorrectAnswerID,
		Difficulty:      uint(q.Difficulty),
		CategoryID:      q.CategoryID,
	}}
	
	return res, nil
}
