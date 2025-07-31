package models

import "time"

type Message struct {
	ID int `json:"id" gorm:"primaryKey"`
	AuthorID int `json:"author_id"`
	Author User `json:"author" gorm:"foreignKey:AuthorID"`
	ChatID int `json:"chat_id"`
	Chat Chat `json:"chat" gorm:"foreignKey:ChatID"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}