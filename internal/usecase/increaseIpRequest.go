package usecase

import "time"

func (u *Usecase) IncreaseIPRequest(ip string, duration time.Duration) (int64, error) {
	total, err := u.redisRepository.IncreaseIPRequest(ip, duration)
	if err != nil {
		return 0, err
	}

	return total, nil
}
