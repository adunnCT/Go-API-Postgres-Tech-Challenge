package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/adunnCT/blog/internal/models"
)

type userCreator interface {
	Create(ctx context.Context, user models.User) (models.User, error)
}

// response represents the response for creating a user.
type createUserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleCreate(logger *slog.Logger, userCreator userCreator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		res, problems, err := decodeValid[models.User](r)

		user := res

		if err != nil {
			for prob := range problems {
				logger.ErrorContext(ctx, "failed to parse req body: ", slog.String("error", problems[prob]))
			}

			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		userRes, err := userCreator.Create(ctx, user)

		if err != nil {
			logger.ErrorContext(ctx, "failed to create user", slog.String("error", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := createUserResponse{
			ID:       userRes.ID,
			Name:     userRes.Name,
			Email:    userRes.Email,
			Password: userRes.Password,
		}

		// Encode the response model as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.ErrorContext(
				ctx,
				"failed to encode response",
				slog.String("error", err.Error()),
			)

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}
