package bookmark

import "github.com/google/uuid"

type CreateBookmarkInput struct {
	UserID uuid.UUID `json:"user_id"`
	URL    string    `json:"url" binding:"required,url"`
	Title  string    `json:"title"`
}
