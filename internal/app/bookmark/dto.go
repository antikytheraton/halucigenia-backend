package bookmark

type CreateBookmarkInput struct {
	UserID string `json:"user_id"`
	URL    string `json:"url" binding:"required,url"`
	Title  string `json:"title"`
}
