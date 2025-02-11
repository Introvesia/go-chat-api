package models

import "time"

type ChannelModel struct {
	ID             string `gorm:"primaryKey"`
	OrganizationID string `gorm:"size:255"`
	Name           string `gorm:"size:255"`
	CreatedDate    time.Time
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
