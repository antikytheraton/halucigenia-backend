package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/internal/platform/config"
	_ "github.com/heroku/x/hmetrics/onload"
)

func Run() int {
	c, err := config.Load(os.Args[1:])
	if err != nil {
		log.Println("error loading config: %v", err)
		return 1
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"OK": "Hello, World!"})
	})

	router.Run(":" + c.HTTP.Port)
	return 0
}
