package bookmark

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, bookmark *Bookmark) error
	FindByID(ctx context.Context, id uuid.UUID) (*Bookmark, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]*Bookmark, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
