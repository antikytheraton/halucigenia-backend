package main

import (
	"log"
	"os"

	"github.com/antikytheraton/halucigenia-backend/cmd/api"
)

func main() {
	log.Println("Server starting...")
	os.Exit(api.Run())
}
