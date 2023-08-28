package questionservice

import (
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) GetQuestion(req param.GetQuestionRequest) (param.GetQuestionResponse, error) {
	const op = "questionservice.GetQuestion"

	q, err := s.repo.GetQuestionByID(req.ID)
	if err != nil {
		return param.GetQuestionResponse{}, richerror.New(op).WithErr(err)
	}

	res := param.GetQuestionResponse{Data: param.Question{
		ID:              q.ID,
		Question:        q.Question,
		PossibleAnswers: s.mapFromPossibleAnswersEntity(q.PossibleAnswers),
		CorrectAnswerID: q.CorrectAnswerID,
		Difficulty:      uint(q.Difficulty),
		CategoryID:      q.CategoryID,
	}}
	return res, nil
}

func (s Service) mapFromPossibleAnswersEntity(q []entity.PossibleAnswer) []param.Answer {
	possibleAnswers := make([]param.Answer, len(q))
	for idx, p := range q {
		possibleAnswers[idx].Text = p.Text
		possibleAnswers[idx].Choice = uint(p.Choice)
	}

	return possibleAnswers
}

func (s Service) mapToPossibleAnswerEntity(q []param.Answer) []entity.PossibleAnswer {
	possibleAnswers := make([]entity.PossibleAnswer, len(q))
	for idx, p := range q {
		possibleAnswers[idx].Text = p.Text
		possibleAnswers[idx].Choice = entity.PossibleAnswerChoice(p.Choice)
	}

	return possibleAnswers
}
