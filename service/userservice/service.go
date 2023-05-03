package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"game-app/entity"
	"game-app/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO - we should verify by phone number by verification code

	// validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is valid")
	}

	// TODO - check the password with regex
	// validate password
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password length must be grater than 8")
	}

	// check uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); !isUnique || err != nil {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	// validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name must more than 3 chars")
	}

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
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return RegisterResponse{User: createdUser}, nil

}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// check existence of phone number in repository

	// get user by phone number

	// compare user.Password with req.Password

	return LoginResponse{}, nil

}

func getMd5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
