package main

import "github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals"

func main() {
	// Create a new Gin router with default middleware
	router := internals.Router()

	// Start the server on port 8080
	router.Run(":8080")
}
