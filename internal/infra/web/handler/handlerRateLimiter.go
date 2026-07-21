package handler

import (
	"net/http"
)

func (h *Handler) RateLimiter(w http.ResponseWriter, r *http.Request) {
	h.usecase.IncreaseIPRequest("192.168.0.1", 60)
	_, _ = w.Write([]byte("Rate Limiter"))
}
