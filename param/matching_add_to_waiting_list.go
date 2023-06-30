package param

import (
	"game-app/entity"
	"time"
)

type AddToWaitingListRequest struct {
	UserID   uint
	Category entity.Category
}

type AddToWaitingListResponse struct {
	Timeout time.Duration `json:"timeout"`
}
