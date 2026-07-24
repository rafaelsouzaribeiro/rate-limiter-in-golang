package usecase

import "time"

func (u *Usecase) Block(ip string, blockDuration time.Duration) error {
	return u.redisRepository.Block(ip, blockDuration)
}
