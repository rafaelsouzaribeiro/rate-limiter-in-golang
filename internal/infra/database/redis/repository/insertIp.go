package repository

import "context"

func (r *Redis) InsertIp(ip string) error {
	ctx := context.Background()

	err := r.client.Set(ctx, "ip", ip, 0).Err()
	if err != nil {
		panic(err)
	}
	return nil
}
