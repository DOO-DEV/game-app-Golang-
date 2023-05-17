package userservice

import (
	"fmt"
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	const op = "userservice.Login"

	// TODO - it would be better to user separate method for existence check and getUserByPhoneNumber
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}

	// compare user.Password with req.Password
	if user.Password != getMd5Hash(req.Password) {
		return param.LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error")
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error")
	}

	return param.LoginResponse{Tokens: param.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, User: param.UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}}, nil

}
