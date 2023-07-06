package redis

import (
	"context"
	"game-app/entity"
	"github.com/labstack/gommon/log"
	"time"
)

func (a Adapter) Publish(event entity.Event, payload string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := a.Client().Publish(ctx, string(event), payload).Err(); err != nil {
		log.Errorf("publish err %v\n", err)
		// TODO - log error
		// TODO - update metrics
	}
}
