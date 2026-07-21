package usecase

func (u *Usecase) IsIPBlocked(ip string) (bool, error) {
	return u.redisRepository.IsIPBlocked(ip)
}
