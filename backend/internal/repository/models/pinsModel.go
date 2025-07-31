package models

type Pin struct {
	ID int `json:"id" gorm:"primaryKey"`
	UserID int `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	Title string `json:"title"`
	Attachment string `json:"attachment"`
}