package usecase

func (u *Usecase) BlockIp(ip string, blockSeconds int) error {
	return u.redisRepository.BlockIP(ip, blockSeconds)
}
