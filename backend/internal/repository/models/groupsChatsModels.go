package models

import "time"

type GroupChat struct {
    ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
    Title      string    `json:"title" gorm:"not null;size:200"`
	Users     []User    `json:"users,omitempty" gorm:"many2many:group_chat_users;"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
    
    // Связи
    Messages []Message `json:"messages,omitempty" gorm:"foreignKey:ChatID"`
}