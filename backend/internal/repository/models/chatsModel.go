package models

type Chat struct {
	ID int `json:"id" gorm:"primaryKey"`
	UserID int `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	Messages []Message `json:"messages" gorm:"foreignKey:ChatID"`
}