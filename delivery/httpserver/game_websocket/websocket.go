package game_websocket

import (
	"context"
	"encoding/json"
	"game-app/adapter/redis"
	"game-app/entity"
	"game-app/logger"
	"game-app/pkg/claims"
	"game-app/pkg/protobufencoder"
	"game-app/pkg/slice"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h Handler) GameWs(c echo.Context) error {
	// get user id
	// subscribe for created game
	// check user id in game events
	// send back the game id to users
	claims := claims.GetClaimsFromEchoContext(c)

	redisAdapter := redis.New(h.redisconfig)
	subscriber := redisAdapter.Client().Subscribe(context.Background(), string(entity.GameCreatedGameEvent))

	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
	if err != nil {
		logger.Logger.Error("websocket connection failed", zap.Error(err))
	}

	defer conn.Close()

	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			logger.Logger.Error("cant consume messages", zap.Error(err))
		}

		payload := protobufencoder.DecodeNewGameGameEvent(msg.Payload)

		if !slice.DoesExist(payload.PlayerIDs, claims.UserID) {
			continue
		}

		response := struct {
			GameID uint `json:"game_id"`
		}{
			GameID: payload.ID,
		}
		byteRes, err := json.Marshal(response)
		if err != nil {
			logger.Logger.Error("cant consume messages", zap.Error(err))
		}

		err = wsutil.WriteServerMessage(conn, ws.OpText, byteRes)
		if err != nil {
			logger.Logger.Error("cant consume messages", zap.Error(err))
		}
	}

}
