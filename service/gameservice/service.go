package gameservice

import (
	"game-app/param"
	"game-app/pkg/richerror"
)

type Repository interface {
	CreateGame(PlayerIDs []uint) (uint, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CreateNewGame(req param.CreateNewGameRequest) (param.CreateNewGameResponse, error) {
	const op = "gameservice.CreateNewGame"

	gameID, err := s.repo.CreateGame(req.PlayerIDs)
	if err != nil {
		return param.CreateNewGameResponse{}, richerror.New(op).WithErr(err)
	}

	return param.CreateNewGameResponse{GameID: gameID}, nil
}
