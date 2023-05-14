package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"game-app/entity"
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

type ProfileRequest struct {
	UserID uint `json:"id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}

func New(auth AuthGenerator, repo Repository) Service {
	return Service{repo: repo, auth: auth}
}

func getMd5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
