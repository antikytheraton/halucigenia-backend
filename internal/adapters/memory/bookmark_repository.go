package memory

import (
	"context"
	"log"
	"sync"

	domain "github.com/antikytheraton/halucigenia-backend/internal/domain/bookmark"
	"github.com/google/uuid"
)

var _ domain.Repository = (*BookmarkRepository)(nil)

type BookmarkRepository struct {
	mu        sync.RWMutex
	bookmarks map[string]*domain.Bookmark
}

func NewBookmarkRepository() *BookmarkRepository {
	return &BookmarkRepository{
		bookmarks: make(map[string]*domain.Bookmark),
	}
}

func (r *BookmarkRepository) Save(ctx context.Context, b *domain.Bookmark) (*domain.Bookmark, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id, _ := uuid.NewRandom()
	b.ID = id
	r.bookmarks[id.String()] = b
	return b, nil
}

func (r *BookmarkRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Bookmark, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	log.Printf("%+v", r.bookmarks)
	b, ok := r.bookmarks[id.String()]
	if !ok {
		return nil, domain.ErrBookmarkNotFound
	}
	return b, nil
}

func (r *BookmarkRepository) ListByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Bookmark, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var bookmarks []*domain.Bookmark
	for _, b := range r.bookmarks {
		bookmarks = append(bookmarks, b)
	}
	return bookmarks, nil
}

func (r *BookmarkRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.bookmarks[id.String()]; !ok {
		return domain.ErrBookmarkNotFound
	}
	delete(r.bookmarks, id.String())
	return nil
}
