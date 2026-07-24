package middleware

import (
	"net"
	"net/http"
	"strings"
)

func (h *Middleware) extractIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return strings.TrimSpace(xrip)
	}

	addr := r.RemoteAddr
	if addr == "" {
		return ""
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		host = addr
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return host
	}

	if ip.IsLoopback() {
		return "127.0.0.1"
	}

	if v4 := ip.To4(); v4 != nil {
		return v4.String()
	}
	return ip.String()
}
