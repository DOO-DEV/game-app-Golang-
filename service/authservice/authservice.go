package authservice

import (
	"game-app/entity"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type Config struct {
	SignKey               string        `koanf:"sign_key"`
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string        `koanf:"access_subject"`
	RefreshSubject        string        `koanf:"refresh_subject"`
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, user.Role, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, user.Role, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func (s Service) createToken(userID uint, role entity.Role, subject string, expDuration time.Duration) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDuration)),
		},
		UserID: userID,
		Role:   role,
	}
	// crate a signer for hs256
	// TODO - replace rs256 with rs256
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
