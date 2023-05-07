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
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"user"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}

type ProfileRequest struct {
	UserID uint `json:"id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}

func New(auth AuthGenerator, repo Repository) Service {
	return Service{repo: repo, auth: auth}
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

	var resp RegisterResponse
	resp.User.ID = createdUser.ID
	resp.User.Name = createdUser.Name
	resp.User.PhoneNumber = createdUser.PhoneNumber
	return resp, nil

}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - it would be better to user separate method for existence check and getUserByPhoneNumber
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	// compare user.Password with req.Password
	if user.Password != getMd5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error")
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error")
	}

	return LoginResponse{Tokens: Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, User: UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}}, nil

}

// all request inputs for intractor/service should be sanitizing

func (s Service) GetProfile(req ProfileRequest) (ProfileResponse, error) {
	// get user by id
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I don'  expected the repository call return "record not found" error,
		//because I assume the interactor input is sanitized
		// TODO - we can use rich error
		return ProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return ProfileResponse{Name: user.Name}, nil
}

func getMd5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
