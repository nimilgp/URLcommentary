package api

import (
	"log/slog"
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/response"
	"github.com/tomasen/realip"
)

func (s *APIServer) logAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw := response.NewMetricsResponseWriter(w)
		next.ServeHTTP(mw, r)

		var (
			ip     = realip.FromRequest(r)
			method = r.Method
			url    = r.URL.String()
			proto  = r.Proto
		)

		userAttrs := slog.Group("user", "ip", ip)
		requestAttrs := slog.Group("request", "method", method, "url", url, "proto", proto)
		responseAttrs := slog.Group("repsonse", "status", mw.StatusCode, "size", mw.BytesCount)

		s.logger.Info("access", userAttrs, requestAttrs, responseAttrs)
	})
}
