package dto

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type UserInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}
