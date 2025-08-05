package main

import (
	"log"

	"github.com/XxThunderBlastxX/thunder-api/internal/router"
	"github.com/XxThunderBlastxX/thunder-api/internal/server"
)

func main() {
	// Initialize the server application
	s := server.New()

	// Register all the application routes
	router.New(s).RegisterRoutes()

	// Start the server and listen on the configured port
	if err := s.Listen(":" + s.Config.Port); err != nil {
		log.Fatal(err)
	}
}
