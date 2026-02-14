package http

import (
	"net/http"

	app "github.com/antikytheraton/halucigenia-backend/internal/app/bookmark"
	domain "github.com/antikytheraton/halucigenia-backend/internal/domain/bookmark"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HandlerConfig struct {
	Service *app.Service
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
	var req app.CreateBookmarkInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Get user ID from auth context
	req.UserID = uuid.MustParse("f2ed0041-71e1-4ac9-9dc6-09537e2bf449")

	err := h.Service.CreateBookmark(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "bookmark created"})
}

// GetBookmarks handles the GET /bookmarks?tag=
func (h *Handler) GetBookmarks(c *gin.Context) {
	// TODO: Get user ID from auth context
	UserID := uuid.MustParse("f2ed0041-71e1-4ac9-9dc6-09537e2bf449")

	bookmarks, err := h.Service.GetBookmarks(c.Request.Context(), UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookmarks": bookmarks})
}

// GetBookmarkByID handles the GET /bookmarks/:id
func (h *Handler) GetBookmarkByID(c *gin.Context) {
	id := c.Param("id")
	bookmarkID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidBookmarkID.Error()})
		return
	}
	bookmark, err := h.Service.GetBookmarkByID(c.Request.Context(), bookmarkID)
	if err == domain.ErrBookmarkNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookmark": bookmark})
}

// DeleteBookmark handles the DELETE /bookmarks/:id
func (h *Handler) DeleteBookmark(c *gin.Context) {
	bookmarkID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidBookmarkID.Error()})
		return
	}
	_ = h.Service.DeleteBookmark(c.Request.Context(), bookmarkID)
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
