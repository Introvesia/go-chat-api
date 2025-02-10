package models

type ChannelMemberModel struct {
	ID          string `gorm:"primaryKey"`
	ChannelID   string `gorm:"size:255"`
	UserID      string `gorm:"size:255"`
	CreatedDate string `gorm:"size:255"`
	UpdatedDate string `gorm:"size:255"`
	DeletedDate string `gorm:"size:255"`
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
