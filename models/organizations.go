package models

import "time"

type OrganizationModel struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"size:255"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time
	DeletedDate time.Time
}

func (OrganizationModel) TableName() string {
	return "organizations"
}

type OrganizationSerializer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	DeletedDate string `json:"deleted_date"`
}
