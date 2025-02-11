package models

import "time"

type UserModel struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	Username        string    `gorm:"size:100;uniqueIndex;not null"`
	Email           string    `gorm:"size:255;uniqueIndex;not null"`
	Password        string    `gorm:"size:100"`
	FirstName       string    `gorm:"size:50;not null"`
	LastName        string    `gorm:"size:50"`
	IsEmailVerified bool      `gorm:"default:false"`
	CreatedDate     time.Time `gorm:"autoCreateTime"`
	UpdatedDate     time.Time
	DeletedDate     time.Time
}

func (UserModel) TableName() string {
	return "users"
}

type UserSerializer struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	IsEmailVerified bool   `json:"is_email_verified"`
	CreatedDate     string `json:"created_date"`
	UpdatedDate     string `json:"updated_date"`
	DeletedDate     string `json:"deleted_date"`
}
