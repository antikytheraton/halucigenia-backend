package http

import (
	"github.com/gin-gonic/gin"
)

// NewRouter creates a new Gin router with the given handler.
func NewRouter(h *Handler) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/", h.HealthCheck)

	v1 := router.Group("/api/v1")
	{
		bookmarkAPI := v1.Group("/bookmarks")
		bookmarkAPI.POST("", h.CreateBookmark)
		bookmarkAPI.GET("", h.GetBookmarks)
		bookmarkAPI.GET("/:id", h.GetBookmarkByID)
		bookmarkAPI.DELETE("/:id", h.DeleteBookmark)

		tagAPI := v1.Group("/tags")
		tagAPI.POST("", h.CreateTag)
		tagAPI.GET("", h.GetTags)
	}
	return router
}
