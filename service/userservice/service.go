package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"game-app/dto"
	"game-app/entity"
	"game-app/pkg/richerror"
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

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	User   dto.UserInfo `json:"user"`
	Tokens Tokens       `json:"tokens"`
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

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	const op = "userservice.Login"

	// TODO - it would be better to user separate method for existence check and getUserByPhoneNumber
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
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
	}, User: dto.UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}}, nil

}

// all request inputs for intractor/service should be sanitizing

func (s Service) GetProfile(req ProfileRequest) (ProfileResponse, error) {
	const op = "userservice.Profile"

	// get user by id
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I don'  expected the repository call return "record not found" error,
		//because I assume the interactor input is sanitized
		// TODO - we can use rich error
		return ProfileResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"req": req})
	}
	return ProfileResponse{Name: user.Name}, nil
}

func getMd5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
