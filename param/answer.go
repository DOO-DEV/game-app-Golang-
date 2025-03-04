package param

type InsertAnswersRequest struct {
	QuestionID uint     `json:"question_id"`
	Data       []Answer `json:"data"`
}

type InsertAnswersResponse struct {
}

type UpdateAnswerRequest struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"question_id"`
	Data       Answer `json:"data"`
}

type UpdateAnswerResponse struct {
	ID   uint   `json:"ID"`
	Data Answer `json:"date"`
}

type DeleteAnswerRequest struct {
	ID uint `json:"ID"`
}

type DeleteAnswerResponse struct {
	Message string `json:"message"`
}

type GetAnswersRequest struct {
	QuestionID uint
}

type GetAnswersResponse struct {
	QuestionID uint
	Data       []Answer
}
