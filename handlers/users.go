package handlers

import (
	"go-chat-api/models"
)

// In-memory user storage
var users = map[string]models.UserModel{
	"admin":    {ID: "1", Username: "admin", Password: "password123"},
	"user1":    {ID: "2", Username: "user1", Password: "securepass"},
	"john_doe": {ID: "3", Username: "john_doe", Password: "mypassword"},
}
