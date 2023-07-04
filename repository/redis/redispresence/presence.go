package redispresence

import (
	"context"
	"game-app/pkg/richerror"
	"time"
)

func (d DB) Upsert(ctx context.Context, key string, timestamp int64, expirationTime time.Duration) error {
	const op = richerror.Op("redispresence.Upsert")

	_, err := d.adapter.Client.Set(ctx, key, timestamp, expirationTime).Result()
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
