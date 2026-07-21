package usecase

import "github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/irepository"

type Usecase struct {
	redisRepository irepository.IRedisRepository
}

func NewUsecase(redisRepository irepository.IRedisRepository) *Usecase {
	return &Usecase{
		redisRepository: redisRepository,
	}
}
