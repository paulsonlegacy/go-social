package model

import (
	"context"
	"database/sql"
)

type Models struct {
	Users interface {
		Create(context.Context, *User) error
	}
	Posts interface {
		Create(context.Context, *Post) error
	}
}

func NewModel(db *sql.DB) Models {
	return Models{
		Users: &UserModel{db},
		Posts: &PostModel{db},
	}
}