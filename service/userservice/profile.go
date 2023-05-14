package userservice

import (
	"game-app/pkg/richerror"
)

func (s Service) GetProfile(req ProfileRequest) (ProfileResponse, error) {
	const op = "userservice.Profile"

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I don'  expected the repository call return "record not found" error,
		//because I assume the interactor input is sanitized
		// TODO - we can use rich error
		return ProfileResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"req": req})
	}
	return ProfileResponse{Name: user.Name}, nil
}
