package handler

import (
	"net/http"
)

func (h *Handler) RateLimiter(w http.ResponseWriter, r *http.Request) {
	h.usecase.InsertIp("192")
	_, _ = w.Write([]byte("Rate Limiter"))
}
