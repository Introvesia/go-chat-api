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

	// Find user on DB
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
	var registry struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&registry); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check username existence
	db := libs.ConnectDB()
	var count int64
	db.Model(&models.UserModel{}).Where("username = ? OR email = ?", registry.Username, registry.Email).Count(&count)

	if count > 0 {
		http.Error(w, "User already exists", http.StatusUnauthorized)
		return
	}

	// Hash password
	hashedPassword, err := hashPassword(registry.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Create new user
	newUser := models.UserModel{
		Username:  registry.Username,
		Email:     registry.Email,
		Password:  hashedPassword,
		FirstName: registry.FirstName,
		LastName:  registry.LastName,
	}

	// Insert user into the database
	if err := db.Create(&newUser).Error; err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(responseSuccess)
}

// Forgot password
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	// TODO: Write forgot password logic
	json.NewEncoder(w).Encode(responseSuccess)
}

// Reset password
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	// TODO: Write reset password logic
	json.NewEncoder(w).Encode(responseSuccess)
}
