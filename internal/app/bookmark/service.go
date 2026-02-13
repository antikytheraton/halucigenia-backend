package bookmark

import (
	"context"

	"github.com/antikytheraton/halucigenia-backend/internal/domain/bookmark"
)

type Service struct {
	repo bookmark.Repository
}

func NewService(repo bookmark.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, input *CreateBookmarkInput) error {
	b := &bookmark.Bookmark{
		UserID: input.UserID,
		URL:    input.URL,
		Title:  input.Title,
	}
	return s.repo.Save(ctx, b)
}
