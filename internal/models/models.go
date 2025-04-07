package models

import (
	"context"
	"database/sql"
)


// Model interface defines common CRUD methods for all models
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