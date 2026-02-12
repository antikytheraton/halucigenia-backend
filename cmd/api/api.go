package api

import (
	"log"
	"net/http"
	"os"

	"github.com/antikytheraton/halucigenia-backend/internal/platform/config"
	"github.com/antikytheraton/halucigenia-backend/internal/platform/db"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func Run() int {
	c, err := config.Load(os.Args[1:])
	if err != nil {
		log.Println("error loading config: %v", err)
		return 1
	}
	sql, err := db.Open(c.Database.URL)
	if err != nil {
		log.Printf("error opening database: %v", err)
		return 1
	}
	defer sql.Close()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"OK": "Hello, World!"})
	})

	router.Run(":" + c.HTTP.Port)
	return 0
}
