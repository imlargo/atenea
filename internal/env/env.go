package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Initialize loads environment variables from .env file
func Initialize() error {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Validate required environment variables
	requiredEnvVars := []string{
		// Add your required env variables here
		// Example: "DATABASE_URL",
		// Example: "API_KEY",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return fmt.Errorf("required environment variable %s is not set", envVar)
		}
	}

	return nil
}
