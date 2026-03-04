package store

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type PostsStore struct {
	db *sql.DB
}

type Post struct {
	ID        int64     `json:"id"`
	Context   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
}
type UpdatePost struct {
	Context *string   `json:"content"`
	Title   *string   `json:"title"`
	Tags    *[]string `json:"tags"`
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

func (s *PostsStore) GetById(ctx context.Context, id int64) (*Post, error) {
	log.Printf("%+v - %+v", ctx, id)
	query := `SELECT id, user_id, title, content, created_at, updated_at, tags
	FROM posts
	WHERE id = $1`

	var post Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Context,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, err

}

func (s *PostsStore) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id=$1`
	result, err := s.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()

	if rows == 0 {
		return ErrNotFound
	}
	// cascase delete
	query_comments := `DELETE FROM comments WHERE post_id=$1`
	result_comments, err := s.db.ExecContext(ctx, query_comments, id)
	if err != nil {
		return err
	}

	rows_comments, err := result_comments.RowsAffected()
	if rows_comments == 0 {
		return ErrNotFound
	}

	return nil
}
func (s *PostsStore) Patch(ctx context.Context, id int64, post *UpdatePost) error {
	query := `
	UPDATE posts
	SET
			title = COALESCE($1, title),
            content = COALESCE($2, content),
            tags    = COALESCE($3, tags)
	WHERE id=$4`

	log.Printf("id=%d post=%+v", id, post)

	var tags interface{}
	if post.Tags != nil {
		tags = pq.Array(post.Tags)
	} else {
		tags = nil
	}
	log.Printf("tags=%v", tags)
	result, err := s.db.ExecContext(ctx, query, post.Title, post.Context, tags, id)
	if err != nil {
		return err
	}

	log.Printf("result: %s", result)

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}
	log.Printf("rows: %d", rows)

	return nil
}
