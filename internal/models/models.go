package models

import (
	"context"
	"database/sql"
)


// Model defines CRUD operations every model should have.
// This allows us to depend on abstractions instead of concrete structs.
// Real models (PostModel, UserModel, etc) implement this interface.
//
// Makes testing easier & keeps handlers decoupled from DB logic.
type Model[ModelType any] interface {
	Create(context.Context, *ModelType) (*ModelType, error)
	GetAll(context.Context) ([]ModelType, error)
	GetByID(context.Context, int64) (*ModelType, error)
	Update(context.Context, *ModelType) (*ModelType, error)
	Delete(context.Context, int64) error
}

// Models struct holds all model instances
type Models struct {
	Users UserModel
	Posts PostModel
	Comments CommentModel
}

// NewModels initializes the models with a shared DB connection
func NewModels(db *sql.DB) Models {
	return Models{
		Users: NewUserModel(db),
		Posts: NewPostModel(db),
		Comments: NewCommentModel(db),
	}
}