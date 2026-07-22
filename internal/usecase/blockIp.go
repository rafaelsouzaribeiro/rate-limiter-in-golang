package usecase

import "time"

func (u *Usecase) BlockIp(ip string, blockDuration time.Duration) error {
	return u.redisRepository.BlockIP(ip, blockDuration)
}
