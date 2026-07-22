package repository

import (
	"context"
	"fmt"
	"time"
)

func (r *Redis) IncreaseIPRequest(ip string, seconds int) (int64, error) {
	ctx := context.Background()
	key := fmt.Sprintf("%s:count", ip)

	total, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if err := r.client.Expire(ctx, key, time.Duration(seconds)*time.Second).Err(); err != nil {
		return 0, err
	}

	return total, nil
}
