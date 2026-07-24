package usecase

func (u *Usecase) IsBlocked(ip string) (bool, error) {
	return u.redisRepository.IsBlocked(ip)
}
