package models

import "time"

type Pin struct {
    ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
    UserID     int       `json:"user_id" gorm:"not null"`
    User       User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Title      string    `json:"title" gorm:"not null;size:200"`
    Content    string    `json:"content" gorm:"type:text"` // Добавляем контент
    CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}