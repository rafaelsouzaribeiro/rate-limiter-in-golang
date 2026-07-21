package repository

import (
	"context"
	"fmt"
	"time"
)

func (r *Redis) BlockIP(ip string, blockSeconds int) error {
	ctx := context.Background()
	key := fmt.Sprintf("%s:block", ip)

	return r.client.Set(ctx, key, 1, time.Duration(blockSeconds)*time.Second).Err()
}
