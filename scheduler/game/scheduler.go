package gamescheduler

import (
	"context"
	"game-app/config"
	"game-app/logger"
	"game-app/param"
	"game-app/service/gameservice"
	"github.com/go-co-op/gocron"
	"time"
)

type Config struct {
	GameMakerIntervalInSeconds int           `koanf:"game_maker_interval_in_seconds"`
	GameMakerTimeoutInSeconds  time.Duration `koanf:"game_maker_timeout"`
}

type Scheduler struct {
	sch     *gocron.Scheduler
	gameSvc gameservice.Service
	config  Config
}

func New(cfg config.Config, gameSvc gameservice.Service) Scheduler {
	return Scheduler{
		sch:     gocron.NewScheduler(time.UTC),
		gameSvc: gameSvc,
		config:  cfg.GameScheduler,
	}
}

func (s Scheduler) Start() {
	s.sch.Every(s.config.GameMakerIntervalInSeconds).Second().Do(s.CreateNewGame)
	s.sch.StartAsync()
}

func (s Scheduler) CreateNewGame() {
	ctx, cancel := context.WithTimeout(context.Background(), s.config.GameMakerTimeoutInSeconds)
	defer cancel()

	if _, err := s.gameSvc.CreateNewGame(ctx, param.CreateNewGameRequest{}); err != nil {
		logger.Logger.Error(err.Error())
	}
}
