package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/heroku/x/hmetrics/onload"

	bookmark_http "github.com/antikytheraton/halucigenia-backend/internal/adapters/http"
	"github.com/antikytheraton/halucigenia-backend/internal/adapters/postgres"
	bookmark_app "github.com/antikytheraton/halucigenia-backend/internal/app/bookmark"
	"github.com/antikytheraton/halucigenia-backend/internal/platform/config"
	"github.com/antikytheraton/halucigenia-backend/internal/platform/db"
)

func Run(args []string) int {
	log.Println("Server starting...")
	c, err := config.Load(args)
	if err != nil {
		log.Println("error loading config: %w", err)
		return 1
	}
	db, err := db.Open(c.Database.URL)
	if err != nil {
		log.Println("error opening database: %w", err)
		return 1
	}

	repo := postgres.NewBookmarkRepository(db)
	service := bookmark_app.NewService(repo)
	handler := bookmark_http.NewHandler(service)
	router := bookmark_http.NewRouter(handler)

	server := &http.Server{
		Addr:         ":" + c.HTTP.Port,
		Handler:      router,
		ReadTimeout:  c.HTTP.ReadTimeout,
		WriteTimeout: c.HTTP.WriteTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	<-ctx.Done()

	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), c.HTTP.GracefulTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Error shutting down server: %v", err)
		return 1
	}
	if err = db.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
		return 1
	}
	return 0
}
