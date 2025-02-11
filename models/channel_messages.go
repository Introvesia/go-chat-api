package models

import "time"

type ChannelMessageModel struct {
	ID          string `gorm:"primaryKey"`
	ChannelID   string `gorm:"size:255"`
	Type        string `gorm:"size:255"`
	Message     string `gorm:"size:255"`
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate time.Time
}

func (ChannelMessageModel) TableName() string {
	return "channel_messages"
}

type ChannelMessageSerializer struct {
	ID          string `json:"id"`
	ChannelID   string `json:"channel_id"`
	Type        string `json:"type"`
	Message     string `json:"message"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	DeletedDate string `json:"deleted_date"`
}
