package usecase

func (u *Usecase) IncreaseIPRequest(ip string, seconds int) (int64, error) {
	total, err := u.redisRepository.IncreaseIPRequest(ip, seconds)
	if err != nil {
		return 0, err
	}

	return total, nil
}
