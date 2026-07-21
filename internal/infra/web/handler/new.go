package handler

import "github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/usecase"

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
