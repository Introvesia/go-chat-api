package handlers

import (
	"encoding/json"
	"go-chat-api/libs"
	"go-chat-api/models"
	"net/http"
)

// Get channels
func GetChannels(w http.ResponseWriter, r *http.Request) {
	db := libs.ConnectDB()
	var channels []models.ChannelModel
	err := db.Model(&models.ChannelModel{}).Find(&channels).Error
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}

// Get channel members
func GetChannelMembers(w http.ResponseWriter, r *http.Request) {
	db := libs.ConnectDB()
	var members []models.ChannelMemberModel
	err := db.Model(&models.ChannelMemberModel{}).Find(&members).Error
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}
