package presenceservice

import (
	"context"
	"fmt"
	"game-app/param"
	"game-app/pkg/richerror"
	"time"
)

type Config struct {
	ExpirationTime time.Duration `koanf:"expiration_time"`
	Prefix         string        `koanf:"prefix"`
}

type Repository interface {
	Upsert(ctx context.Context, key string, timestamp int64, expirationTime time.Duration) error
}

type Service struct {
	repo   Repository
	config Config
}

func New(repo Repository, config Config) Service {
	return Service{
		repo:   repo,
		config: config,
	}
}

func (s Service) UpsertPresence(ctx context.Context, req param.UpsertPresenceRequest) (param.UpsertPresenceResponse, error) {
	const op = "precenseservice.UpsertPresence"

	if err := s.repo.Upsert(ctx, fmt.Sprintf("%s:%d", s.config.Prefix, req.UserID), req.Timestamp, s.config.ExpirationTime); err != nil {
		return param.UpsertPresenceResponse{}, richerror.New(op).WithErr(err)
	}

	return param.UpsertPresenceResponse{}, nil
}

func (s Service) GetPresence(ctx context.Context, req param.GetPresenceRequest) (param.GetPresenceResponse, error) {
	fmt.Println("req", req)
	// TODO - implement me
	return param.GetPresenceResponse{Items: []param.GetPresenceItem{
		{UserID: 1, Timestamp: 1234567}, {UserID: 2, Timestamp: 74397418},
	}}, nil
}
