package bookmark

import (
	"context"

	"github.com/antikytheraton/halucigenia-backend/internal/domain/bookmark"
	"github.com/google/uuid"
)

type Service struct {
	bookmarkRepo bookmark.Repository
}

func NewService(repo bookmark.Repository) *Service {
	return &Service{
		bookmarkRepo: repo,
	}
}

// CreateBookmark creates a new bookmark for the user.
func (s *Service) CreateBookmark(ctx context.Context, input *CreateBookmarkInput) error {
	b := &bookmark.Bookmark{
		UserID: input.UserID,
		URL:    input.URL,
		Title:  input.Title,
	}
	return s.bookmarkRepo.Save(ctx, b)
}

// GetBookmarkByID retrieves a bookmark by its ID.
func (s *Service) GetBookmarkByID(ctx context.Context, id uuid.UUID) (*bookmark.Bookmark, error) {
	return s.bookmarkRepo.FindByID(ctx, id)
}

// GetBookmarks retrieves all bookmarks for a given user.
func (s *Service) GetBookmarks(ctx context.Context, userID uuid.UUID) ([]*bookmark.Bookmark, error) {
	return s.bookmarkRepo.ListByUserID(ctx, userID)
}

// DeleteBookmark deletes a bookmark by its ID.
func (s *Service) DeleteBookmark(ctx context.Context, id uuid.UUID) error {
	return s.bookmarkRepo.Delete(ctx, id)
}
