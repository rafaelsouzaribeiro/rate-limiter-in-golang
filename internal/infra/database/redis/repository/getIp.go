package repository

import "context"

func (r *Redis) GetIp() (string, error) {
	ctx := context.Background()

	ip, err := r.client.Get(ctx, "ip").Result()
	if err != nil {
		panic(err)
	}
	return ip, nil
}
