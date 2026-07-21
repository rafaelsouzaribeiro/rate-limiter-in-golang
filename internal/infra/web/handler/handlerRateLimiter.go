package handler

import (
	"net/http"
)

func (h *Handler) RateLimiter(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("RateLimiter"))
}
