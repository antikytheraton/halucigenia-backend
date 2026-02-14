package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/antikytheraton/halucigenia-backend/internal/domain/bookmark"
	"github.com/google/uuid"
)

var _ bookmark.Repository = (*BookmarkRepository)(nil)

type BookmarkRepository struct {
	db *sql.DB
}

func NewBookmarkRepository(db *sql.DB) *BookmarkRepository {
	return &BookmarkRepository{db: db}
}

// Save inserts a new bookmark into the database.
func (r *BookmarkRepository) Save(ctx context.Context, b *bookmark.Bookmark) (*bookmark.Bookmark, error) {
	var row bookmark.Bookmark
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO bookmarks (user_id, url, title)
	  VALUES ($1, $2, $3)
		RETURNING id, user_id, url, title, created_at, updated_at`,
		b.UserID, b.URL, b.Title,
	).Scan(&row.ID, &row.UserID, &row.URL, &row.Title, &row.CreatedAt, &row.UpdatedAt)
	if err != nil {
		log.Printf("failed to save bookmark: %v", err)
		return nil, bookmark.ErrBookmarkSaveFailed
	}
	return &row, nil
}

// FindByID retrieves a bookmark by its ID.
func (r *BookmarkRepository) FindByID(ctx context.Context, id uuid.UUID) (*bookmark.Bookmark, error) {
	var b bookmark.Bookmark
	err := r.db.QueryRowContext(ctx,
		`
		SELECT id, user_id, url, title, created_at, updated_at
		FROM bookmarks
		WHERE id = $1
		`,
		id.String(),
	).Scan(&b.ID, &b.UserID, &b.URL, &b.Title, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, bookmark.ErrBookmarkNotFound
		}
		return nil, err
	}
	return &b, nil
}

// ListByUserID retrieves all bookmarks for a given user.
func (r *BookmarkRepository) ListByUserID(ctx context.Context, userID uuid.UUID) ([]*bookmark.Bookmark, error) {

	var bookmarks []*bookmark.Bookmark
	rows, err := r.db.QueryContext(ctx,
		`
		SELECT id, user_id, url, title, created_at, updated_at
		FROM bookmarks
		WHERE user_id = $1
		`,
		userID.String(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b bookmark.Bookmark
		rowErr := rows.Scan(&b.ID, &b.UserID, &b.URL, &b.Title, &b.CreatedAt, &b.UpdatedAt)
		if rowErr != nil {
			log.Printf("failed to scan bookmark: %v", rowErr)
			return nil, nil
		}
		bookmarks = append(bookmarks, &b)
	}

	return bookmarks, nil
}

// Delete removes a bookmark by its ID.
func (r *BookmarkRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx,
		`DELETE FROM bookmarks WHERE id = $1`,
		id.String(),
	)
	return err
}
