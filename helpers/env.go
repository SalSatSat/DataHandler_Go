package helpers

import (
	"log"
	"os"

	// Import godotenv
	"github.com/joho/godotenv"
)

// Use godot package to load/read the .env file and
// Return the value of the key
func EnvVariable(key string) string {

	// Load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
