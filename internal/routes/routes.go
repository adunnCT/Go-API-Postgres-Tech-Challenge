package routes

import (
	"log/slog"
	"net/http"

	"github.com/adunnCT/blog/internal/handlers"
	"github.com/adunnCT/blog/internal/services"
)

func AddRoutes(mux *http.ServeMux, logger *slog.Logger, usersService *services.UsersService) {
	// Read a user
	mux.Handle("GET /api/users/{id}", handlers.HandleRead(logger))
}
