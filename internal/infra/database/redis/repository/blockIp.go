package repository

import (
	"context"
	"fmt"
	"time"
)

func (r *Redis) Block(ip string, blockDuration time.Duration) error {
	ctx := context.Background()
	key := fmt.Sprintf("%s:block", ip)

	return r.client.Set(ctx, key, 1, blockDuration).Err()
}
