package models

import "time"

type ChannelModel struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	OrganizationID uint
	Name           string    `gorm:"size:255"`
	CreatedDate    time.Time `gorm:"autoCreateTime"`
	UpdatedDate    time.Time
	DeletedDate    time.Time
}

func (ChannelModel) TableName() string {
	return "channels"
}

type ChannelSerializer struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	CreatedDate    string `json:"created_date"`
	UpdatedDate    string `json:"updated_date"`
	DeletedDate    string `json:"deleted_date"`
}
