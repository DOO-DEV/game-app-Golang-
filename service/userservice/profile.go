package userservice

import (
	"context"
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) GetProfile(ctx context.Context, req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.Profile"

	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		// I don'  expected the repository call return "record not found" error,
		//because I assume the interactor input is sanitized
		// TODO - we can use rich error
		return param.ProfileResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"req": req})
	}
	return param.ProfileResponse{Name: user.Name}, nil
}
