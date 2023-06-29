package config

import (
	"game-app/service/authservice"
	"time"
)

func Default() Config {
	return Config{
		Debug: true,
		Auth: authservice.Config{
			SignKey:               "jwt_secret",
			AccessSubject:         "at",
			RefreshSubject:        "rt",
			AccessExpirationTime:  time.Hour * 24,
			RefreshExpirationTime: time.Hour * 24 * 7,
		},
	}
}
