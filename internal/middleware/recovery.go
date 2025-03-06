package middleware

import (
	"log/slog"
	"net/http"
)

func Recover(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {

					// log the error
					logger.ErrorContext(r.Context(), "Recovered from panic", slog.String("error", err.(string)))

					// return 500 status code to the client
					http.Error(w, "Server error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
