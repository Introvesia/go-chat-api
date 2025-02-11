package handlers

import (
	"encoding/json"
	"go-chat-api/libs"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ResponseSuccess struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

var responseSuccess ResponseSuccess

func init() {
	responseSuccess = ResponseSuccess{Code: 200, Status: "success"}
}

// Generate JWT Token
func Login(w http.ResponseWriter, r *http.Request) {
	// Get JWT secret and convert it to bytes
	jwtSecret := []byte(libs.GetEnv("JWT_SECRET"))

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

// Register
func Register(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseSuccess)
}

// Forgot password
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseSuccess)
}

// Reset password
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseSuccess)
}
