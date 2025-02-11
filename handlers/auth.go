package handlers

import (
	"encoding/json"
	"fmt"
	"go-chat-api/libs"
	"go-chat-api/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ResponseSuccess struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

var responseSuccess ResponseSuccess

func init() {
	responseSuccess = ResponseSuccess{Code: 200, Status: "success"}
}

// Hash password using bcrypt
func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// Compare hashed password
func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
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

	db := libs.ConnectDB()
	var user models.UserModel
	err := db.Where("username = ?", creds.Username).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Validate credentials by directly accessing the map
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	fmt.Print(passwordHash)
	if checkPasswordHash(creds.Password, passwordHash) {
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
