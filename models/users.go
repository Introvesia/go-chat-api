package models

type UserModel struct {
	ID              string `gorm:"primaryKey"`
	Username        string `gorm:"size:255"`
	Password        string `gorm:"size:255"`
	FirstName       string `gorm:"size:255"`
	LastName        string `gorm:"size:255"`
	IsEmailVerified bool   `gorm:"default:false"`
	CreatedDate     string `gorm:"size:255"`
	UpdatedDate     string `gorm:"size:255"`
	DeletedDate     string `gorm:"size:255"`
}

func (UserModel) TableName() string {
	return "users"
}

type UserSerializer struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	IsEmailVerified bool   `json:"is_email_verified"`
	CreatedDate     string `json:"created_date"`
	UpdatedDate     string `json:"updated_date"`
	DeletedDate     string `json:"deleted_date"`
}
