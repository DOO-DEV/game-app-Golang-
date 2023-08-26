package param

type Answer struct {
	Text   string `json:"text"`
	Choice uint   `json:"choice"`
}

type CreateNewQuestionRequest struct {
	Question        string   `json:"question"`
	PossibleAnswers []Answer `json:"possible_answers"`
	CorrectAnswerID uint     `json:"correct_answer_id"`
	Difficulty      uint     `json:"difficulty"`
	CategoryID      uint     `json:"category_id"`
}

type CreateNewQuestionResponse struct {
	Question        string   `json:"question"`
	PossibleAnswers []Answer `json:"possible_answers"`
	CorrectAnswerID uint     `json:"correct_answer_id"`
	Difficulty      uint     `json:"difficulty"`
	CategoryID      uint     `json:"category_id"`
}

type UpdateQuestionResponse struct {
}

type UpdateQuestionRequest struct {
	Question        string   `json:"question"`
	PossibleAnswers []Answer `json:"possible_answers"`
	CorrectAnswerID uint     `json:"correct_answer_id"`
	Difficulty      uint     `json:"difficulty"`
	CategoryID      uint     `json:"category_id"`
}

type DeleteQuestionRequest struct {
	ID uint `json:"id"`
}

type DeleteQuestionResponse struct {
	Message string `json:"message"`
}

type GetQuestionResponse struct {
}

type GetQuestionRequest struct {
	ID uint `json:"id"`
}

type GetQuestionsByCategoryRequest struct {
	CategoryID uint `json:"category_id"`
}

type GetQuestionsByCategoryResponse struct {
}
