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
	GetPresence(ctx context.Context, prefixKey string, userIDs []uint) (map[uint]int64, error)
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
	list, err := s.repo.GetPresence(ctx, s.config.Prefix, req.UserIDs)
	if err != nil {
		return param.GetPresenceResponse{}, err
	}

	resp := param.GetPresenceResponse{}
	for k, v := range list {
		resp.Items = append(resp.Items, param.GetPresenceItem{
			UserID:    k,
			Timestamp: v,
		})
	}

	return resp, nil
}
