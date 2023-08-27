package questionservice

import (
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) UpdateQuestion(req param.UpdateQuestionRequest) (param.UpdateQuestionResponse, error) {
	const op = "questionservice.UpdateQuestion"

	q := entity.Question{
		Question:        req.Data.Question,
		PossibleAnswers: s.mapToPossibleAnswerEntity(req.Data.PossibleAnswers),
		CorrectAnswerID: req.Data.CorrectAnswerID,
		Difficulty:      entity.QuestionDifficulty(req.Data.Difficulty),
		CategoryID:      req.Data.CategoryID,
	}
	question, err := s.repo.UpdateQuestion(q)
	if err != nil {
		return param.UpdateQuestionResponse{}, richerror.New(op).WithErr(err)
	}
	res := param.UpdateQuestionResponse{Data: param.Question{
		ID:              question.ID,
		Question:        question.Question,
		PossibleAnswers: s.mapFromPossibleAnswersEntity(question.PossibleAnswers),
		CorrectAnswerID: question.CorrectAnswerID,
		Difficulty:      uint(question.Difficulty),
		CategoryID:      question.CategoryID,
	}}

	return res, nil
}
