package connection

import "github.com/redis/go-redis/v9"

func NewConnection(addr, password string) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Password: password,
		},
	)

}
