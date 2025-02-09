package handlers

import (
	"basic-golang-api/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// In-memory user storage
var users = map[string]models.User{
	"admin":    {ID: "1", Username: "admin", Password: "password123"},
	"user1":    {ID: "2", Username: "user1", Password: "securepass"},
	"john_doe": {ID: "3", Username: "john_doe", Password: "mypassword"},
}

// Load environment variables
func loadEnv() {
	_ = godotenv.Load()
}

// Generate JWT Token
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	loadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate credentials by directly accessing the map
	user, exists := users[creds.Username]
	if !exists || user.Password != creds.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
