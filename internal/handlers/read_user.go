package handlers

import (
	"log/slog"
	"net/http"
)

func HandleRead(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the status code to 200 OK
		w.WriteHeader(http.StatusOK)

		id := r.PathValue("id")

		if id == "" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		// Write the response body, simply echo the ID back out
		_, err := w.Write([]byte(id))
		if err != nil {
			// Handle error if response writing fails
			logger.ErrorContext(r.Context(), "failed to write response", slog.String("error", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
