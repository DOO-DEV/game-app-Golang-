package userservice

import (
	"fmt"
	"game-app/dto"
	"game-app/entity"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	// create new user
	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		// TODO - replace md5 with bcyrpt
		Password: getMd5Hash(req.Password),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	var resp dto.RegisterResponse
	resp.User.ID = createdUser.ID
	resp.User.Name = createdUser.Name
	resp.User.PhoneNumber = createdUser.PhoneNumber
	return resp, nil

}
