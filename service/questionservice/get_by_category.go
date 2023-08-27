package questionservice

import (
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) GetQuestionsByCategory(req param.GetQuestionsByCategoryRequest) (param.GetQuestionsByCategoryResponse, error) {
	const op = "questionservice.GEtQuestionsByCategory"

	questions, err := s.repo.GetQuestionsByCategory(req.CategoryID)
	if err != nil {
		return param.GetQuestionsByCategoryResponse{}, richerror.New(op).WithErr(err)
	}

	var res param.GetQuestionsByCategoryResponse
	for idx, q := range questions {
		res.Data[idx].Question = q.Question
		res.Data[idx].ID = q.ID
		res.Data[idx].CategoryID = q.CategoryID
		res.Data[idx].Difficulty = uint(q.Difficulty)
		res.Data[idx].PossibleAnswers = s.mapFromPossibleAnswersEntity(q.PossibleAnswers)
		res.Data[idx].CorrectAnswerID = q.CorrectAnswerID
	}

	return res, nil
}
