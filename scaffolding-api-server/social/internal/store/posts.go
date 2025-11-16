package store

import (
	"context"
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type PostsStore struct {
	db *sql.DB
}

type Post struct {
	ID        int64    `json:"id"`
	Context   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	log.Printf("ctx: %+v %+v", ctx, post)
	query := `INSERT INTO posts(content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	row := s.db.QueryRowContext(ctx, query, post.Context, post.Title, post.UserID, pq.Array(post.Tags))
	err := row.Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}
