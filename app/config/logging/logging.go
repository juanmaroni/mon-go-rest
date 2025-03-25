package logging

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

var Logger *slog.Logger

func LoadJSONLogger() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func HttpRequestLogging(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			logger.Info("HTTP Request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.Duration("duration", time.Since(start)),
			)

			next.ServeHTTP(w, r)
		})
	}
}
