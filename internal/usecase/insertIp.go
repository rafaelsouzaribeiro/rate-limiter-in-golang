package usecase

func (u *Usecase) InsertIp(ip string) error {
	err := u.redisRepository.InsertIp(ip)
	if err != nil {
		return err
	}
	return nil
}
