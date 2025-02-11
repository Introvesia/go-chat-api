package libs

import (
	"go-chat-api/models"
)

func AutoMigrate() {
	LoadEnv()
	db := ConnectDB()
	db.AutoMigrate(&models.UserModel{})
	db.AutoMigrate(&models.OrganizationModel{})
	db.AutoMigrate(&models.ChannelModel{})
	db.AutoMigrate(&models.ChannelMessageModel{})
	db.AutoMigrate(&models.ChannelMemberModel{})
}
