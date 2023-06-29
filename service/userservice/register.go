package userservice

import (
	"fmt"
	"game-app/entity"
	"game-app/param"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	// create new user
	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		// TODO - replace md5 with bcyrpt
		Password: getMd5Hash(req.Password),
		Role:     entity.UserRole,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	var resp param.RegisterResponse
	resp.User.ID = createdUser.ID
	resp.User.Name = createdUser.Name
	resp.User.PhoneNumber = createdUser.PhoneNumber
	return resp, nil

}
