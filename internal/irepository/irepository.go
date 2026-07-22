package irepository

import "time"

type IRedisRepository interface {
	InsertIp(ip string) error
	IncreaseIPRequest(ip string, duration time.Duration) (int64, error)
	IsIPBlocked(ip string) (bool, error)
	BlockIP(ip string, blockDuration time.Duration) error
}
