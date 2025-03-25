package models

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"encoding/json"
)

// Post structure
type Post struct {
	ID        int64    `json:"id"`
	UserID    int64    `json:"user_id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// PostModel implements the Model interface for Post
type PostModel struct {
	db *sql.DB
}

// NewPostModel initiates a new PostModel
func NewPostModel(db *sql.DB) PostModel {
	return PostModel{db: db}
}

// Create inserts a new post into the database
func (postmodel PostModel) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (user_id, title, content, tags, created_at, updated_at)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`

	log.Println("Before conversion:", *post)

	// Convert `[]string` to JSON string
	tagsJSON, err := json.Marshal(post.Tags)
	if err != nil {
		return err
	}

	log.Println("Converted Tags JSON:", string(tagsJSON))

	result, err := postmodel.db.ExecContext(
		ctx,
		query,
		post.UserID,
		post.Title,
		post.Content,
		tagsJSON,
	)

	if err != nil {
		return err
	}

	// Retrieve the last inserted ID
	post.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	// Set created_at and updated_at manually if needed
	post.CreatedAt = "NOW()"
	post.UpdatedAt = "NOW()"

	return nil
}

// GetByID retrieves a post by its ID
func (postmodel PostModel) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `SELECT id, user_id, title, content, tags, created_at, updated_at FROM posts WHERE id = ? LIMIT 1`
	row := postmodel.db.QueryRowContext(ctx, query, id)


	var post Post
	var tags string
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &tags, &post.CreatedAt, &post.UpdatedAt)
	
	if err == sql.ErrNoRows {

		return nil, nil // Return nil instead of an error if no post is found

	}else if err != nil {

		return nil, err // If error occured

	}

	// Trim spaces when splitting tags
	post.Tags = strings.FieldsFunc(tags, func(r rune) bool {
		return r == ',' || r == ' '
	})

	return &post, nil
}

// Update modifies an existing post
func (postmodel PostModel) Update(ctx context.Context, post *Post) error {
	query := `UPDATE posts SET title = ?, content = ?, tags = ?, updated_at = Now() WHERE id = ? LIMIT 1`
	_, err := postmodel.db.ExecContext(
		ctx,
		query,
		post.Title, 
		post.Content, 
		strings.Join(post.Tags, ","), 
		post.ID,
	)
	return err
}

// Delete removes a post by its ID
func (postmodel PostModel) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id = ? LIMIT 1`
	_, err := postmodel.db.ExecContext(ctx, query, id)
	return err
}