package matchign

import (
	"context"
	"fmt"
	"game-app/entity"
	"game-app/pkg/richerror"
	"game-app/pkg/timestamp"
	"github.com/redis/go-redis/v9"
	"strconv"
)

const WaitingListPrefix = "waitinglist"

func (d DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = "matching.AddToWaitingList"

	ctx := context.Background()

	key := fmt.Sprintf("%s:%s", WaitingListPrefix, category)
	idStr := strconv.Itoa(int(userID))

	_, err := d.adapter.Client.ZAdd(ctx, key, redis.Z{
		Score:  float64(timestamp.Now()),
		Member: idStr,
	}).Result()

	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil

}
