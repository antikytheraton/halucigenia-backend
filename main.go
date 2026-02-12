package main

import (
	"log"
	"os"

	"github.com/heroku/go-getting-started/cmd/api"
)

func main() {
	log.Println("Server starting...")
	os.Exit(api.Run())
}
