package services

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/adunnCT/blog/internal/models"
)

// UsersService is a service capable of performing CRUD operations for
// models.User models.
type UsersService struct {
	logger *slog.Logger
	db     *sql.DB
}

// NewUsersService creates a new UsersService and returns a pointer to it.
func NewUsersService(logger *slog.Logger, db *sql.DB) *UsersService {
	return &UsersService{logger: logger, db: db}
}

// Create attempts to create the provided user, returning a fully hydrated
// models.User or an error

func (s *UsersService) Create(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

// Read attempts to read a user from the database using the provided id. A
// fully hydrated models.User or error is returned
func (s *UsersService) Read(ctx context.Context, id uint64) (models.User, error) {
	return models.User{}, nil
}

// Update attempts to perform an update of the user with the provided id,
// updating it to reflect the properties on the provided patch object. A
// models.User or an error is returned
func (s *UsersService) Update(ctx context.Context, id uint64, patch models.User) (models.User, error) {
	return models.User{}, nil
}

// Delete attempts to delete the user with the provided id. An error is
// returned if the delete fails
func (s *UsersService) Delete(ctx context.Context, id uint64) error {
	return nil
}

// List attempts to list all users in the database. A slice of models.User
// or an error is returned
func (s *UsersService) List(ctx context.Context, id uint64) ([]models.User, error) {
	return []models.User{}, nil
}
