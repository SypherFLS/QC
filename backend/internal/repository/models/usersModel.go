package models

import (
	"errors"
	"time"
)

type User struct {
    ID       int       `json:"id" gorm:"primaryKey;autoIncrement"`
    Username string    `json:"username" gorm:"unique;not null;size:50"`
    Email    string    `json:"email" gorm:"unique;not null;size:100"`
    Password string    `json:"-" gorm:"not null;size:255"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
    
    Pins  []Pin  `json:"pins,omitempty" gorm:"foreignKey:UserID"`
    Chats []Chat `json:"chats,omitempty" gorm:"foreignKey:UserID"`
    GroupChats []GroupChat `json:"group_chats,omitempty" gorm:"many2many:group_chat_users;"`
}

func (u *User) Validate() error {
    if u.Username == "" {
        return errors.New("username is required")
    }
    if u.Email == "" {
        return errors.New("email is required")
    }
    return nil
}