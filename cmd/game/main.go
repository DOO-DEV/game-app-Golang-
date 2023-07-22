package game

import (
	"context"
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/entity"
	"game-app/logger"
	"game-app/param"
	"game-app/pkg/protobufencoder"
	redisgame "game-app/repository/redis/game"
	"game-app/service/gameservice"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main(cfg config.Config) {
	redisAdapter := redis.New(cfg.Redis)
	gameRepository := redisgame.New(redisAdapter)
	gameSvc := gameservice.New(gameRepository)

	topic := entity.MatchingUsersMatchedEvent
	subscriber := redisAdapter.Client().Subscribe(context.Background(), string(topic))

	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			logger.Logger.Error("cant consume messages", zap.Error(err))
		}

		payload := protobufencoder.DecodeMatchingUsersMatchedEvent(msg.Payload)

		res, err := gameSvc.CreateNewGame(param.CreateNewGameRequest{PlayerIDs: payload.UserIDs})
		if err != nil {
			logger.Logger.Error("can't create a new game", zap.Error(err))
		}

		pbGame := protobufencoder.EncodeNewGameGameEvent(res.GameID)
		redisAdapter.Publish(entity.GameCreatedGameEvent, pbGame)
	}
}

func New(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "game",
		Short: "run game microservice for create game between users",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}
}
