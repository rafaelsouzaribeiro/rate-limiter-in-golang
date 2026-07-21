package middleware

import (
	"net"
	"net/http"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func RateLimitByIP(rps rate.Limit, burst int, ttl time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//ip := extractIP(r)

			/*mu.Lock()
			c, ok := clients[ip]
			if !ok {
				c = &clientLimiter{
					limiter: rate.NewLimiter(rps, burst),
				}
				clients[ip] = c
			}
			c.lastSeen = time.Now()
			allowed := c.limiter.Allow()
			mu.Unlock()

			if !allowed {
				http.Error(w, "too many requests", http.StatusTooManyRequests)
				return
			}*/

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
