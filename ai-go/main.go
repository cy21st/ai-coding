package main

import (
	"log"

	"ai-go/config"
	"ai-go/database"
	"ai-go/routes"
)

func main() {
	// Initialize environment
	config.InitEnv()

	// Initialize database connections
	database.InitDB()
	database.InitRedis()

	// Setup router
	r := routes.SetupRouter()

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
