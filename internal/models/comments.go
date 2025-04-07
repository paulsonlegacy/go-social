package models

import (
	"time"
	"context"
	"database/sql"
)

// TYPES

// Comment structure
type Comment struct {
	ID			int64 `json:"id"`
	UserID		int64 `json:"user_id"`
	PostID		int64 `json:"post_id"`
	ParentID    int64 `json:"parent_id"`
	Content		string `json:"content"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// CommentModel implements the Model interface for Comment
type CommentModel struct {
	db *sql.DB
}


// FUNCTIONS

// NewCommentModel initiates a new CommentModel
func NewCommentModel(db *sql.DB) CommentModel {
	return CommentModel{db: db}
}





// Create inserts a new comment into the database
func (commentmodel CommentModel) Create(ctx context.Context, comment *Comment) (*Comment, error) {
	
	now := time.Now().Format("2006-01-02 15:04:05") // MySQL DATETIME format

	query := `
		INSERT INTO comments 
		(user_id, post_id, parent_id, content)
		VALUES (?, ?, ?, ?)
	`

	result, err := commentmodel.db.ExecContext(
		ctx,
		query,
		comment.UserID,
		comment.PostID,
		comment.ParentID,
		comment.Content,
	)

	if err != nil {

		return nil, err

	}

	// Retrieving and appending the last inserted ID to struct object
	comment.ID, err = result.LastInsertId()

	if err != nil {

		return nil, err

	}

	// Set created_at and updated_at manually if needed
	comment.CreatedAt = now
	comment.UpdatedAt = now

	return comment, nil
}





// GetByID retrieves a comment by its ID
func (commentmodel CommentModel) GetByID(ctx context.Context, id int64) (*Comment, error) {
	query := `
		SELECT 
		id, user_id, title, content, tags, created_at, updated_at 
		FROM comments 
		WHERE id = ? 
		LIMIT 1
	`

	comment := Comment{}
	row := commentmodel.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&comment.ID, 
		&comment.UserID, 
		&comment.PostID,
		&comment.ParentID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)

	if err == sql.ErrNoRows {

		return nil, nil // Return nil instead of an error if no post is found

	} else if err != nil {

		return nil, err // If error occured

	}

	return &comment, nil
}





// Update modifies an existing comment
func (commentmodel CommentModel) Update(ctx context.Context, comment *Comment) (*Comment, error) {
	
	now := time.Now().Format("2006-01-02 15:04:05") // MySQL DATETIME format
	
	query := `
		UPDATE comments 
		SET content = ?, updated_at = ?
		WHERE id = ? 
		LIMIT 1
	`

	_, err := commentmodel.db.ExecContext(
		ctx,
		query,
		comment.Content,
		now,
		comment.ID,
	)

	// Set created_at and updated_at manually if needed
	comment.UpdatedAt = now

	return comment, err
}





// Delete removes a comment by its ID
func (commentmodel CommentModel) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM comments
		WHERE id = ?
		LIMIT 1
	`

	_, err := commentmodel.db.ExecContext(ctx, query, id)

	return err
}