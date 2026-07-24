package middleware

import (
	"net/http"

	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/pkg/duration"
	"github.com/spf13/viper"
)

const limitMessage = "you have reached the maximum number of requests or actions allowed within a certain time frame"

func (h *Middleware) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		item := r.Header.Get("API_KEY")
		total, err := h.CheckToken(item, w, r)
		if err != nil {
			return
		}

		item = h.extractIP(r)
		total, err = h.CheckIp(item, w, r)
		if err != nil {
			return
		}

		if total > int64(viper.GetInt("MAX_REQUEST")) {
			_ = h.usecase.Block(item, duration.GetDuration(viper.GetString("BLOCK_TIME")))
			http.Error(w, limitMessage, http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
