package redis

import (
	"context"
	"game-app/entity"
	"game-app/logger"
	"go.uber.org/zap"
	"time"
)

func (a Adapter) Publish(event entity.Event, payload string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := a.Client().Publish(ctx, string(event), payload).Err(); err != nil {
		logger.Logger.Error("publish", zap.Error(err))
		// TODO - update metrics
	}
}
