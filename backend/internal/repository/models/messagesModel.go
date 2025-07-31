package models

import "time"

type Message struct {
    ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
    AuthorID  int       `json:"author_id" gorm:"not null"`
    Author    User      `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
    ChatID    int       `json:"chat_id" gorm:"not null"`
    Chat      Chat      `json:"chat,omitempty" gorm:"foreignKey:ChatID"`
    Content   string    `json:"content" gorm:"not null;type:text"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}