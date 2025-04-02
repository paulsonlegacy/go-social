package models

import (
	"context"
	"database/sql"
)


// User structure
type User struct {
	ID int64          `json:"id"`
	FirstName string  `json:"first_name"`
	LastName string  `json:"last_name"`
	Email string      `json:"email"`
	Username string   `json:"username"`
	Password string   `json:"-"`
	IsAdmin bool      `json:"is_admin"`
	IsActive bool     `json:"is_active"`
	JoinedAt string  `json:"joined_at"`
}

// UserModel implements the Model interface for User
type UserModel struct {
	db *sql.DB
}

// NewUserModel initiates a new UserModel
func NewUserModel(db *sql.DB) UserModel {
	return UserModel{db: db}
}

// Creates a new user
func (usermodel UserModel) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (first_name, last_name, username, email, password)
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, joined_at
	`

	err := usermodel.db.QueryRowContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.JoinedAt,
	)

	if err != nil {
		return err
	}
	return nil
}




// GetAll retrieves all users from the database.
// It executes a SQL query to fetch user details, maps the result to a slice of user structs,
// and returns the list of users or an error if any occurs.
//
// Parameters:
// - ctx (context.Context): The context for managing request deadlines and cancellations.
//
// Returns:
// - ([]User): A slice of user structs containing all retrieved users.
// - (error): An error if the database query or data processing fails.
func (usermodel UserModel) GetAll(ctx context.Context) ([]User, error) {
	query := `
		SELECT 
		first_name, last_name, username, email, joined_at
		FROM users
	`
	rows, err := usermodel.db.QueryContext(ctx, query)

	// Return error if query execution fails
	if err != nil {
		return nil, err
	}
	// Ensure rows are closed after function execution to prevent resource leaks
	defer rows.Close()

	// slice to hold each user struct
	var users []User

	// Iterate over each row in the result set
	for rows.Next() {
		var user User       // For parsing each user result into

		// Scan row values into the user struct fields
		err := rows.Scan(
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.JoinedAt,
		)

		// Return error if scanning fails
		if err != nil {
			return nil, err
		}

		// Append the parsed user to the list
		users = append(users, user)
	}

	// Return the retrieved users
	return users, nil
}





// GetByID retrieves a user by its ID
func (usermodel UserModel) GetByID(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT 
		first_name, last_name, username, email, joined_at 
		FROM users 
		WHERE id = ? 
		LIMIT 1
	`
	row := usermodel.db.QueryRowContext(ctx, query, id)

	var user User       // For parsing user result into
	
	err := row.Scan(
		&user.FirstName, 
		&user.LastName, 
		&user.Username, 
		&user.Email, 
		&user.JoinedAt,
	)

	if err == sql.ErrNoRows {

		return nil, nil // Return nil instead of an error if no user is found

	} else if err != nil {

		return nil, err // If error occured

	}

	return &user, nil
}





// Update modifies an existing user
func (usermodel UserModel) Update(ctx context.Context, user *User) error {
	query := `
		UPDATE users 
		SET first_name = ?, last_name = ?, username = ?, email = ?
		WHERE id = ? 
		LIMIT 1
	`

	_, err := usermodel.db.ExecContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
	)

	return err
}





// Delete removes a user by its ID
func (usermodel UserModel) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM users 
		WHERE id = ? 
		LIMIT 1
	`
	_, err := usermodel.db.ExecContext(ctx, query, id)

	return err
}
