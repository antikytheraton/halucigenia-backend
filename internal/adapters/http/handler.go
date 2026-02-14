package http

import (
	"net/http"

	"github.com/antikytheraton/halucigenia-backend/internal/app/bookmark"
	"github.com/gin-gonic/gin"
)

type HandlerConfig struct {
	Service *bookmark.Service
	Env     string
}

type Handler struct {
	HandlerConfig
}

// NewHandler creates a new Handler with the given bookmark service.
func NewHandler(c HandlerConfig) *Handler {
	return &Handler{c}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "halucigenia-backend",
		"status":  "ok",
		"env":     h.Env,
	})
}

// CreateBookmark handles the POST /bookmarks
func (h *Handler) CreateBookmark(c *gin.Context) {
	var req bookmark.CreateBookmarkInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Get user ID from auth context
	req.UserID = "f2ed0041-71e1-4ac9-9dc6-09537e2bf449"

	err := h.Service.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "bookmark created"})
}

// GetBookmarks handles the GET /bookmarks?tag=
func (h *Handler) GetBookmarks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"bookmarks": []string{}})
}

// DeleteBookmark handles the DELETE /bookmarks/:id
func (h *Handler) DeleteBookmark(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "bookmark deleted"})
}

// CreateTag handles the POST /tags
func (h *Handler) CreateTag(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "tag created"})
}

// GetTags handles the GET /tags
func (h *Handler) GetTags(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tags": []string{}})
}
