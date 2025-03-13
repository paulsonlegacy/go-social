package model

import (
	"context"
	"database/sql"
)

type User struct {
	ID int64          `json:"id"`
	FirstName string  `json:"first_name"`
	LastName string  `json:"last_name"`
	Email string      `json:"email"`
	Username string   `json:"username"`
	Password string   `json:"-"`
	IsAdmin bool      `json:"is_admin"`
	IsActive bool     `json:"is_active"`
	CreatedAt string  `json:"created_at"`
}

type UserModel struct {
	db *sql.DB
}

func (usermodel *UserModel) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (first_name, last_name, email, username, password, is_admin, is_active , created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at
`

	err := usermodel.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Username,
		user.Password,
		user.IsAdmin,
		user.IsActive,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}