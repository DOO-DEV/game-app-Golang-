package param

type InsertAnswersRequest struct {
	QuestionID uint     `json:"question_id"`
	Data       []Answer `json:"data"`
}

type InsertAnswersResponse struct {
	Message string `json:"message"`
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
