package bookmark

import "context"

type Repository interface {
	Save(ctx context.Context, bookmark *Bookmark) error
	FindByID(ctx context.Context, id string) (*Bookmark, error)
	ListByUserID(ctx context.Context, userID string) ([]*Bookmark, error)
}
