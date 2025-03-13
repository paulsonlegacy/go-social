package model

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// Post contains the post structure
type Post struct {
	ID int64          `json:"id"`
	UserID int64      `json:"user_id"`
	Title string	  `json:"title"`
	Content string    `json:"content"`
	Tags []string     `json:"tags"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// Post Model connects to the database
type PostModel struct {
	db *sql.DB
}

func (postmodel *PostModel) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (user_id, title, content, tags, created_at, updated_at)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := postmodel.db.QueryRowContext(
		ctx,
		query,
		post.UserID,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}