package main

import (
	"os"

	"github.com/antikytheraton/halucigenia-backend/cmd/api"
)

func main() {
	os.Exit(api.Run(os.Args[1:]))
}
