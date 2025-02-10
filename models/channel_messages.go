package models

type ChannelMessageModel struct {
	ID          string `gorm:"primaryKey"`
	ChannelID   string `gorm:"size:255"`
	Type        string `gorm:"size:255"`
	Message     string `gorm:"size:255"`
	CreatedDate string `gorm:"size:255"`
	UpdatedDate string `gorm:"size:255"`
	DeletedDate string `gorm:"size:255"`
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
