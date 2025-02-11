package models

import "time"

type ChannelMemberModel struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	ChannelID   uint
	UserID      string    `gorm:"size:255"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time
	DeletedDate time.Time
}

func (ChannelMemberModel) TableName() string {
	return "channel_members"
}

type ChannelMemberSerializer struct {
	ID          string `json:"id"`
	ChannelID   string `json:"channel_id"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	DeletedDate string `json:"deleted_date"`
}
