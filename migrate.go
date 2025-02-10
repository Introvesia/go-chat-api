package main

import (
	"go-chat-api/libs"
	"go-chat-api/models"
)

func main() {
	libs.LoadEnv()
	db := libs.ConnectDB()
	db.AutoMigrate(&models.UserModel{})
	db.AutoMigrate(&models.OrganizationModel{})
	db.AutoMigrate(&models.ChannelModel{})
	db.AutoMigrate(&models.ChannelMessageModel{})
	db.AutoMigrate(&models.ChannelMemberModel{})
}
