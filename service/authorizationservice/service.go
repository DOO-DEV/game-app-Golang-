package authorizationservice

import (
	"game-app/entity"
	"game-app/pkg/richerror"
)

type Repository interface {
	GetUserPermissionsTitle(userID uint, role entity.Role) ([]entity.PermissionTitle, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CheckAccess(userID uint, role entity.Role, permissions ...entity.PermissionTitle) (bool, error) {
	// check access
	const op = "authorizationservice.CheckAccess"

	permissionTitles, err := s.repo.GetUserPermissionsTitle(userID, role)

	if err != nil {
		return false, richerror.New(op).WithErr(err)
	}

	for _, pt := range permissionTitles {
		for _, p := range permissions {
			if pt == p {
				return true, nil
			}
		}
	}
	return false, nil
}
