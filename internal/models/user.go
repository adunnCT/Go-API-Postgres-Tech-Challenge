package models

import "context"

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

func (u User) Valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	if u.Name == "" {
		problems["name"] = "Name is required"
	}

	if u.Email == "" {
		problems["email"] = "Email is required"
	}

	if u.Password == "" {
		problems["password"] = "Password is required"
	}

	return problems
}
