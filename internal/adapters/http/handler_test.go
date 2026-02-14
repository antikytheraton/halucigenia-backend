package http_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	bookmark_http "github.com/antikytheraton/halucigenia-backend/internal/adapters/http"
	"github.com/antikytheraton/halucigenia-backend/internal/adapters/memory"
	bookmark_app "github.com/antikytheraton/halucigenia-backend/internal/app/bookmark"
	domain "github.com/antikytheraton/halucigenia-backend/internal/domain/bookmark"
)

func TestCreateBookmark(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := memory.NewBookmarkRepository()
	service := bookmark_app.NewService(repo)
	handler := bookmark_http.NewHandler(bookmark_http.HandlerConfig{
		Service: service,
		Env:     "testing",
	})
	router := bookmark_http.NewRouter(handler)

	body := []byte(`{"title":"Example","url":"https://example.com"}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/bookmarks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFindBookmark(t *testing.T) {
	repo := memory.NewBookmarkRepository()
	service := bookmark_app.NewService(repo)
	handler := bookmark_http.NewHandler(bookmark_http.HandlerConfig{
		Service: service,
		Env:     "testing",
	})
	router := bookmark_http.NewRouter(handler)

	// First, create a bookmark to ensure there is something to find
	bookmark, _ := repo.Save(context.TODO(), &domain.Bookmark{
		Title:  "Example",
		URL:    "https://example.com",
		UserID: uuid.MustParse("f2ed0041-71e1-4ac9-9dc6-09537e2bf449"),
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/bookmarks/"+bookmark.ID.String(), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
