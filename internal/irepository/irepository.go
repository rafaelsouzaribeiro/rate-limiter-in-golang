package irepository

import "time"

type IRedisRepository interface {
	IncreaseRequest(ip string, duration time.Duration) (int64, error)
	IsBlocked(ip string) (bool, error)
	Block(ip string, blockDuration time.Duration) error
}
