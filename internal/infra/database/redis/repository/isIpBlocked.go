package repository

import (
	"context"
	"fmt"
)

func (r *Redis) IsBlocked(ip string) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("%s:block", ip)

	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return exists == 1, nil
}
