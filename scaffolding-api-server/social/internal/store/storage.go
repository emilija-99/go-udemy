package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Resource Not Found.")
)

type Storage struct {
	Posts interface {
		GetById(context.Context, int64) (*Post, error)
		Create(context.Context, *Post) error
		Delete(context.Context, int64) error
		Patch(context.Context, int64, *UpdatePost) error
	}
	User interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostsStore{db},
		User:     &UserStore{db},
		Comments: &CommentStore{db},
	}
}
