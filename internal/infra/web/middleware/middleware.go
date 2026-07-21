package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

const limitMessage = "you have reached the maximum number of requests or actions allowed within a certain time frame"

func (h *Middleware) RateLimitByIP() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := extractIP(r)

			blocked, err := h.usecase.IsIPBlocked(ip)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			if blocked {
				http.Error(w, limitMessage, http.StatusTooManyRequests)
				return
			}

			total, err := h.usecase.IncreaseIPRequest(ip, viper.GetInt("TIME_LIMIT")) // TIME_LIMIT
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			if total > int64(viper.GetInt("MAX_REQUEST")) { // MAX_REQUEST
				_ = h.usecase.BlockIp(ip, viper.GetInt("BLOCK_TIME")) // BLOCK_TIME
				http.Error(w, limitMessage, http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func extractIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return strings.TrimSpace(xrip)
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
