package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)


// TYPES

// Post structure
type Post struct {
	ID        int64    `json:"id"`
	UserID    int64    `json:"user_id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Comments  []Comment
}

// PostModel implements the Model interface for Post
type PostModel struct {
	db *sql.DB
}

// FUNCTIONS

// NewPostModel initiates a new PostModel
func NewPostModel(db *sql.DB) PostModel {
	return PostModel{db: db}
}





// Create inserts a new post into the database
func (postmodel PostModel) Create(ctx context.Context, post *Post) (*Post, error) {

	now := time.Now().Format("2006-01-02 15:04:05") // MySQL DATETIME format

	query := `
		INSERT INTO posts (user_id, title, content, tags, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	// Convert `[]string` to JSON string
	tagsJSON, err := json.Marshal(post.Tags)
	if err != nil {
		return nil, err
	}

	result, err := postmodel.db.ExecContext(
		ctx,
		query,
		post.UserID,
		post.Title,
		post.Content,
		tagsJSON,
		now,
		now,
	)

	if err != nil {

		return nil, err

	}

	// Retrieve the last inserted ID
	post.ID, err = result.LastInsertId()

	if err != nil {

		return nil, err

	}

	// Set created_at and updated_at manually if needed
	post.CreatedAt = now
	post.UpdatedAt = now

	return post, nil
}





// GetAll retrieves all posts from the database.
// It executes a SQL query to fetch post details, maps the result to a slice of Post structs,
// and returns the list of posts or an error if any occurs.
//
// Parameters:
// - ctx (context.Context): The context for managing request deadlines and cancellations.
//
// Returns:
// - ([]Post): A slice of Post structs containing all retrieved posts.
// - (error): An error if the database query or data processing fails.
func (postmodel PostModel) GetAll(ctx context.Context) ([]Post, error) {
	query := `
		SELECT 
		id, user_id, title, content, tags, created_at, updated_at 
		FROM posts
	`
	rows, err := postmodel.db.QueryContext(ctx, query)

	// Return error if query execution fails
	if err != nil {
		return nil, err
	}
	// Ensure rows are closed after function execution to prevent resource leaks
	defer rows.Close()

	// slice to hold each post struct
	var posts []Post

	// Iterate over each row in the result set
	for rows.Next() {
		var post Post       // For parsing each post result into
		var tagsJSON string // Holds JSON string representation of tags

		// Scan row values into the Post struct fields
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&tagsJSON, // Tags are stored as a JSON string in the database
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		// Return error if scanning fails
		if err != nil {

			return nil, err
			
		}

		// Convert the JSON string of current post tags
		// into a slice of strings then update post.Tags
		err = json.Unmarshal([]byte(tagsJSON), &post.Tags)

		if err != nil {

			return nil, err

		}

		// Append the parsed post to the list
		posts = append(posts, post)
	}

	// Return the retrieved posts
	return posts, nil
}





// GetByID retrieves a post by its ID
func (postmodel PostModel) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT 
		id, user_id, title, content, tags, created_at, updated_at FROM posts 
		WHERE id = ? 
		LIMIT 1
	`
	row := postmodel.db.QueryRowContext(ctx, query, id)

	var post Post       // For parsing post result into
	var tagsJSON string // Holds JSON string representation of tags
	
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &tagsJSON, &post.CreatedAt, &post.UpdatedAt)

	if err == sql.ErrNoRows {

		return nil, nil // Return nil instead of an error if no post is found

	} else if err != nil {

		return nil, err // If error occured

	}

	// Convert the JSON string of post tags
	// into a slice of strings then update post.Tags
	err = json.Unmarshal([]byte(tagsJSON), &post.Tags)
	
	if err != nil {

		return nil, err

	}

	return &post, nil
}





// Update modifies an existing post
func (postmodel PostModel) Update(ctx context.Context, post *Post) (*Post, error) {
	query := `
		UPDATE posts 
		SET user_id = ?, title = ?, content = ?, tags = ?, updated_at = NOW() 
		WHERE id = ? 
		LIMIT 1
	`

	// Convert `[]string` to JSON string
	tagsJSON, err := json.Marshal(post.Tags)
	if err != nil {
		return nil, err
	}

	_, err = postmodel.db.ExecContext(
		ctx,
		query,
		post.UserID,
		post.Title,
		post.Content,
		tagsJSON,
		post.ID,
	)

	// Set created_at and updated_at manually if needed
	post.UpdatedAt = time.Now().Format("2006-01-02 15:04:05") // MySQL DATETIME format

	return post, err
}





// Delete removes a post by its ID
func (postmodel PostModel) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM posts 
		WHERE id = ? 
		LIMIT 1
	`
	_, err := postmodel.db.ExecContext(ctx, query, id)

	return err
}
