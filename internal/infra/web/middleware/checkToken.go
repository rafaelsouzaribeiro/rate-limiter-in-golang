package middleware

import (
	"errors"
	"net/http"

	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/pkg/duration"
	"github.com/spf13/viper"
)

func (h *Middleware) CheckToken(token string, w http.ResponseWriter, r *http.Request) (int64, error) {

	blocked, err := h.usecase.IsBlocked(token)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return 0, err
	}
	if blocked {
		http.Error(w, limitMessage, http.StatusTooManyRequests)
		return 0, errors.New(limitMessage)
	}

	total, err := h.usecase.IncreaseRequest(token, duration.GetDuration(viper.GetString("TOKEN_LIMIT")))
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return 0, err
	}
	return total, nil
}
