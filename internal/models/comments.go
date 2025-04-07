package models

import (
	"time"
	"context"
	"database/sql"
)

// TYPES

// Comment structure
type Comment struct {
	ID			int64 	`json:"id"`
	UserID		int64 	`json:"user_id"`
	PostID		int64 	`json:"post_id"`
	ParentID    int64 	`json:"parent_id"`
	Content		string 	`json:"content"`
	CreatedAt 	string 	`json:"created_at"`
	UpdatedAt 	string  `json:"updated_at"`
	Commenter   User   	`json:"commenter"`  // Nested struct for commenter details
	Replies  []Comment  `json:"replies"`
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
			c.id, c.user_id, c.title, c.content, c.tags, c.created_at, c.updated_at,
			u.first_name, u.last_name, u.username 
		FROM 
			comments c
		JOIN 
			users u ON u.id = c.user_id
		WHERE 
			c.id = ? 
		LIMIT 1
	`

	comment := Comment{} 	// Comment placeholder
	commenter := User{}		// Commenter placeholder
	row := commentmodel.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&comment.ID, 
		&comment.UserID, 
		&comment.PostID,
		&comment.ParentID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&commenter.FirstName,
		&commenter.LastName,
		&commenter.Username,
	)

	if err == sql.ErrNoRows {

		return nil, nil // Return nil instead of an error if no post is found

	} else if err != nil {

		return nil, err // If error occured

	}

	// Setting commenter
	comment.Commenter = commenter

	repliesQuery := `
		SELECT 
			c.id, c.user_id, c.title, c.content, c.tags, c.created_at, c.updated_at,
			u.first_name, u.last_name, u.username 
		FROM 
			comments c
		JOIN 
			users u ON u.id = c.user_id
		WHERE 
			c.parent_id = ? 
		ORDER BY 
			created_at 
		ASC
	`

	// Querying the DB
	rows, err := commentmodel.db.QueryContext(ctx, repliesQuery, comment.ParentID)

	if err != nil {

		return nil, err

	}

	// Preventing further enumeration 
	defer rows.Close()

	//  Comment slice to hold reply comments structs
	var replies []Comment

	// Looping through comments
	for rows.Next() {

		// Current comment
		var reply Comment

		err := rows.Scan(
			&reply.ID,
			&reply.UserID,
			&reply.PostID,
			&reply.ParentID,
			&reply.Content,
			&reply.CreatedAt,
			&reply.UpdatedAt,
		)

		if err != nil {

			return nil, err

		}

		// Append reply to replies slice
		replies = append(replies, reply)
	}
	
	// Setting comment replies
	comment.Replies = replies
		
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