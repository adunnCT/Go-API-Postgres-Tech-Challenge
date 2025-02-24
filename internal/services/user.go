package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	s.logger.DebugContext(ctx, "Creating user", "id", user.Name)

	res, err := s.db.ExecContext(ctx,
		`INSERT INTO users (name, email, password) VALUES ($1::text, $2::text, $3::text)`,
		user.Name, user.Email, user.Password,
	)

	if err != nil {
		return models.User{}, fmt.Errorf(
			"[in services.UsersService.Create] failed to create user: %w",
			err,
		)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return models.User{}, fmt.Errorf(
			"[in services.UsersService.Create] failed to get created user ID: %w",
			err,
		)
	}

	user.ID = uint(id)

	return user, nil
}

// Read attempts to read a user from the database using the provided id. A
// fully hydrated models.User or error is returned
func (s *UsersService) Read(ctx context.Context, id uint64) (models.User, error) {
	s.logger.DebugContext(ctx, "Reading user", "id", id)

	row := s.db.QueryRowContext(
		ctx,
		`SELECT id, name, email, password FROM users WHERE id = $1::int`,
		id,
	)

	var user models.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.User{}, nil
		default:
			return models.User{}, fmt.Errorf(
				"[in services.UsersService.Read] failed to read user: %w",
				err,
			)
		}
	}

	return user, nil
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
