package repository

import "github.com/redis/go-redis/v9"

type Redis struct {
	client *redis.Client
}

func NewRepository(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}
