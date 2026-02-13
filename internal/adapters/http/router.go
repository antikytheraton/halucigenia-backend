package http

import (
	"github.com/gin-gonic/gin"
)

// NewRouter creates a new Gin router with the given handler.
func NewRouter(h *Handler) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/bookmarks", h.CreateBookmark)
	}
	return router
}
