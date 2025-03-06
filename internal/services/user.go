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
	s.logger.DebugContext(ctx, "Creating user", "name", user.Name)

	err := s.db.QueryRowContext(ctx,
		`INSERT INTO users (name, email, password) VALUES ($1::text, $2::text, $3::text) RETURNING id`,
		user.Name, user.Email, user.Password,
	).Scan(&user.ID)

	if err != nil {
		return models.User{}, fmt.Errorf(
			"[in services.UsersService.Create] failed to create user: %w",
			err,
		)
	}

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
	s.logger.DebugContext(ctx, "Updating user", "id", id)

	res, err := s.db.ExecContext(
		ctx,
		`UPDATE users SET name = $1::text, email = $2::text, password = $3::text WHERE id = $4::int`,
		patch.Name, patch.Email, patch.Password, id,
	)

	if err != nil {
		return models.User{}, fmt.Errorf(
			"[in services.UsersService.Update] failed to update user: %w",
			err,
		)
	}

	numberOfRowsAffected, err := res.RowsAffected()

	if err != nil {
		return models.User{}, fmt.Errorf(
			"[in services.UsersService.Update] failed to get number of rows affected: %w",
			err,
		)
	}

	if numberOfRowsAffected == 0 {
		return models.User{}, fmt.Errorf("[in services.UsersService.Update] no rows were updated")
	}

	return patch, nil
}

// Delete attempts to delete the user with the provided id. An error is
// returned if the delete fails
func (s *UsersService) Delete(ctx context.Context, id uint64) error {
	s.logger.DebugContext(ctx, "Deleting user", "id", id)

	_, err := s.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1::int`, id)

	if err != nil {
		return fmt.Errorf("[in services.UsersService.Delete] failed to delete user: %w", err)
	}

	return nil
}

// List attempts to list all users in the database. A slice of models.User
// or an error is returned
func (s *UsersService) List(ctx context.Context, id uint64) (userList []models.User, err error) {
	s.logger.DebugContext(ctx, "Listing users")

	rows, err := s.db.QueryContext(
		ctx,
		`SELECT * FROM users`,
	)

	if err != nil {
		return []models.User{}, fmt.Errorf("[in services.UsersService.List] failed to get users: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
			return userList, fmt.Errorf("[in services.UsersService.List] failed to scan row: %w", err)
		}

		userList = append(userList, user)
	}

	return
}
