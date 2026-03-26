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
	Context   string    `json:"context"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
	Version   int64     `json:"version"`
}
type UpdatePost struct {
	Context *string   `json:"context"`
	Title   *string   `json:"title"`
	Tags    *[]string `json:"tags"`
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	log.Printf("ctx: %+v %+v", ctx, post)
	query := `INSERT INTO posts(context, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

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
	query := `SELECT id, user_id, title, context, created_at, updated_at, tags, version
	FROM posts
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var post Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Context,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
		&post.Version,
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
	log.Printf("Id: %d", id)
	query := `DELETE FROM posts WHERE id=$1`
	result, err := s.db.ExecContext(ctx, query, id)

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

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
func (s *PostsStore) Patch(ctx context.Context, post *Post) error {
	log.Printf("post: %+v ctx: %+v", post, ctx)
	query := `
	UPDATE posts
	SET title = $1,
	    context = $2,
	    tags = $3,
	    version = version + 1
	WHERE id = $4
	RETURNING version
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var tags interface{}

	if post.Tags != nil {
		tags = pq.Array(post.Tags)
	} else {
		tags = nil
	}

	log.Printf("tags=%v", tags)

	err := s.db.QueryRowContext(ctx, query,
		post.Title,
		post.Context,
		pq.Array(post.Tags),
		post.ID,
	).Scan(&post.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}
	return nil
}
