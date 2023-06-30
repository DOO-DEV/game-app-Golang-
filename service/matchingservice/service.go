package matchingservice

import (
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
	"time"
)

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}

type Repository interface {
	AddToWaitingList(userID uint, category entity.Category) error
}

type Service struct {
	config Config
	repo   Repository
}

func New(repo Repository, config Config) Service {
	return Service{
		config: config,
		repo:   repo,
	}
}

func (s Service) AddToWaitingList(req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {
	const op = "matchingservice.AddToWaitiongList"

	// add user to the waiting list for the given category if not exist

	if err := s.repo.AddToWaitingList(req.UserID, req.Category); err != nil {
		return param.AddToWaitingListResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return param.AddToWaitingListResponse{Timeout: s.config.WaitingTimeout}, nil
}
