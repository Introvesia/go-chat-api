package models

type OrganizationModel struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	CreatedDate string `gorm:"size:255"`
	UpdatedDate string `gorm:"size:255"`
	DeletedDate string `gorm:"size:255"`
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
