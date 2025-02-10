package libs

import (
	"os"

	"github.com/joho/godotenv"
)

// Load environment variables
func LoadEnv() {
	_ = godotenv.Load()
}

// Get environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}
