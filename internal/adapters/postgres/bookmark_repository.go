package postgres

import (
	"context"
	"database/sql"

	"github.com/antikytheraton/halucigenia-backend/internal/domain/bookmark"
)

var _ bookmark.Repository = (*BookmarkRepository)(nil)

type BookmarkRepository struct {
	db *sql.DB
}

func NewBookmarkRepository(db *sql.DB) *BookmarkRepository {
	return &BookmarkRepository{db: db}
}

func (r *BookmarkRepository) Save(ctx context.Context, b *bookmark.Bookmark) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO bookmarks (user_id, url, title)
	  VALUES ($1, $2, $3)`,
		b.UserID, b.URL, b.Title,
	)
	return err
}

func (r *BookmarkRepository) FindByID(ctx context.Context, id string) (*bookmark.Bookmark, error) {
	return nil, nil
}

func (r *BookmarkRepository) ListByUserID(ctx context.Context, userID string) ([]*bookmark.Bookmark, error) {
	return nil, nil
}
