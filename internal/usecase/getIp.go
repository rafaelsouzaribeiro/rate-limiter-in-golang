package usecase

func (u *Usecase) GetIp() (string, error) {
	ip, err := u.redisRepository.GetIp()
	if err != nil {
		return "", err
	}
	return ip, nil
}
