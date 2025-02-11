package models

import "time"

type ChannelMemberModel struct {
	ID          string `gorm:"primaryKey"`
	ChannelID   string `gorm:"size:255"`
	UserID      string `gorm:"size:255"`
	CreatedDate time.Time
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
