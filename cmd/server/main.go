package main

import (
	"ecommerce-backend/internal/routers"
	"log"
)

func main() {
	r := routers.NewRouter()

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
