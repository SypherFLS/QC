package models

type User struct {
	ID int `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Pins []Pin `json:"pins" gorm:"foreignKey:UserID"`
	Chats []Chat `json:"chats" gorm:"foreignKey:UserID"`
}

