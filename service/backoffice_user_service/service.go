package backoffice_user_service

import "game-app/entity"

type Service struct {
}

func New() Service {
	return Service{}
}

func (s Service) ListAllUsers() ([]entity.User, error) {
	// TODO - implement me

	users := make([]entity.User, 0)

	users = append(users, entity.User{
		ID:          0,
		Name:        "fake",
		PhoneNumber: "fake",
		Password:    "fake",
		Role:        entity.AdminRole,
	})

	return users, nil
}
