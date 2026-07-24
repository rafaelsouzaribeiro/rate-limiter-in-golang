package usecase

import "time"

func (u *Usecase) IncreaseRequest(ip string, duration time.Duration) (int64, error) {
	total, err := u.redisRepository.IncreaseRequest(ip, duration)
	if err != nil {
		return 0, err
	}

	return total, nil
}
