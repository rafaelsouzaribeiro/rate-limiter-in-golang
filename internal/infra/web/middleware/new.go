package middleware

import "github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/usecase"

type Middleware struct {
	usecase *usecase.Usecase
}

func NewMiddleware(usecase *usecase.Usecase) *Middleware {
	return &Middleware{
		usecase: usecase,
	}
}
