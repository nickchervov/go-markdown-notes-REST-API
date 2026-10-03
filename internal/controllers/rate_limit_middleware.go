package connectors

import (
	"net"
	"net/http"
	"strconv"

	"github.com/nickchervov/go-markdown-notes-REST-API/pkg/render"
)

func (h *NotesHandler) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "rate limit internal error", http.StatusInternalServerError)
			return
		}

		retryAfter, ok, err := h.svc.RateLimit(r.Context(), 20, ip)
		if err != nil {
			http.Error(w, "rate limit internal error", http.StatusInternalServerError)
			return
		}

		if !ok {
			w.Header().Set("Retry-After", strconv.Itoa(retryAfter))
			render.JSON(w, http.StatusTooManyRequests, map[string]string{"code": "too many requests", "message": "Rate Limited. Try later.", "retry-after": strconv.Itoa(retryAfter)})
			return
		}

		next.ServeHTTP(w, r)
	})
}
