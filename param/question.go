package param

type Answer struct {
	Text   string `json:"text"`
	Choice uint   `json:"choice"`
}

type Question struct {
	ID              uint     `json:"id"`
	Question        string   `json:"question"`
	PossibleAnswers []Answer `json:"possible_answers"`
	CorrectAnswerID uint     `json:"correct_answer_id"`
	Difficulty      uint     `json:"difficulty"`
	CategoryID      uint     `json:"category_id"`
}

type CreateNewQuestionRequest struct {
	Data Question `json:"data"`
}

type CreateNewQuestionResponse struct {
	Data Question `json:"data"`
}

type UpdateQuestionResponse struct {
	Data Question `json:"data"`
}

type UpdateQuestionRequest struct {
	Data Question `json:"data"`
}

type DeleteQuestionRequest struct {
	ID uint `json:"id"`
}

type DeleteQuestionResponse struct {
	Message string `json:"message"`
}

type GetQuestionResponse struct {
	Data Question `json:"data"`
}

type GetQuestionRequest struct {
	ID uint `json:"id"`
}

type GetQuestionsByCategoryRequest struct {
	CategoryID uint `json:"category_id"`
}

type GetQuestionsByCategoryResponse struct {
	Data []Question `json:"data"`
}
