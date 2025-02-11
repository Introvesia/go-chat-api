package models

import "time"

type ChannelMessageModel struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	ChannelID   uint      `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	Type        string    `gorm:"size:255;not null"`
	Message     string    `gorm:"size:255;not null"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
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
