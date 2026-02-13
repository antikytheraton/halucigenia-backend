package http

import (
	"net/http"

	"github.com/antikytheraton/halucigenia-backend/internal/app/bookmark"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *bookmark.Service
}

// NewHandler creates a new Handler with the given bookmark service.
func NewHandler(service *bookmark.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateBookmark(c *gin.Context) {
	var req bookmark.CreateBookmarkInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Get user ID from auth context
	req.UserID = "f2ed0041-71e1-4ac9-9dc6-09537e2bf449"

	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "bookmark created"})
}
