package questionservice

import (
	"context"
	"game-app/entity"
	"game-app/logger"
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) UpdateQuestion(ctx context.Context, req param.UpdateQuestionRequest) (param.UpdateQuestionResponse, error) {
	const op = "questionservice.UpdateQuestion"

	q := entity.Question{
		ID:              req.Data.ID,
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

	if _, err := s.answerClient.InsertAnswers(ctx, param.InsertAnswersRequest{
		QuestionID: req.Data.ID,
		Data:       req.Data.PossibleAnswers,
	}); err != nil {
		logger.Logger.Error(err.Error())
	}

	res := param.UpdateQuestionResponse{Data: param.Question{
		ID:              question.ID,
		Question:        question.Question,
		PossibleAnswers: req.Data.PossibleAnswers,
		CorrectAnswerID: question.CorrectAnswerID,
		Difficulty:      uint(question.Difficulty),
		CategoryID:      question.CategoryID,
	}}

	return res, nil
}
