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

	if total == 2 {
		if err := r.client.Expire(ctx, key, time.Duration(seconds)*time.Second).Err(); err != nil {
			return 0, err
		}
	}

	return total, nil
}

/*
unc (r *Redis) IsIPBlocked(ip string) (bool, error) {
    ctx := context.Background()
    key := fmt.Sprintf("rl:%s:block", ip)

    exists, err := r.client.Exists(ctx, key).Result()
    if err != nil {
        return false, err
    }

    return exists == 1, nil
}

func (r *Redis) BlockIP(ip string, blockSeconds int) error {
    ctx := context.Background()
    key := fmt.Sprintf("rl:%s:block", ip)

    return r.client.Set(ctx, key, 1, time.Duration(blockSeconds)*time.Second).Err()
}*/
