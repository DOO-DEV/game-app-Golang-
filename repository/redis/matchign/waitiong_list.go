package matchign

import (
	"context"
	"fmt"
	"game-app/entity"
	"game-app/pkg/richerror"
	"game-app/pkg/timestamp"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const WaitingListPrefix = "waitinglist"

func (d DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = "matching.AddToWaitingList"

	ctx := context.Background()

	idStr := strconv.Itoa(int(userID))
	_, err := d.adapter.Client().ZAdd(ctx, getCategoryKey(category), redis.Z{
		Score:  float64(timestamp.Now()),
		Member: idStr,
	}).Result()

	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil

}

func (d DB) GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error) {
	const op = "redismatching.GetWaitingListByCategory"

	min := fmt.Sprintf("%d", timestamp.Add(-2*time.Hour))
	max := fmt.Sprintf("%d", timestamp.Now())

	list, err := d.adapter.Client().ZRangeByScoreWithScores(ctx, getCategoryKey(category), &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: 0,
		Count:  0,
	}).Result()

	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	result := make([]entity.WaitingMember, 0)

	for _, l := range list {
		userID, _ := strconv.Atoi(l.Member.(string))

		result = append(result, entity.WaitingMember{
			UserID:    uint(userID),
			Timestamp: int64(l.Score),
			Category:  category,
		})
	}

	return result, nil
}

func getCategoryKey(category entity.Category) string {
	return fmt.Sprintf("%s:%s", WaitingListPrefix, category)
}
