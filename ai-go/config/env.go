package config

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var Env string

// InitEnv initializes the environment configuration
func InitEnv() {
	// Define command line flags
	flag.StringVar(&Env, "env", "dev", "Environment (dev/prod)")
	flag.Parse()

	// Set Gin mode based on environment
	switch Env {
	case "dev":
		gin.SetMode(gin.DebugMode)
		log.Println("Running in development mode")
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		log.Println("Running in production mode")
	default:
		gin.SetMode(gin.DebugMode)
		log.Printf("Unknown environment '%s', defaulting to development mode\n", Env)
	}
}
