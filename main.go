package main

import (
	"fmt"
	"os"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/router"
)

func main() {
	router := router.NewRouter().
		Setup()

	// Start the server on port 8080
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
