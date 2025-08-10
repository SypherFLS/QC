package rInterfaces

import (
	"qc/internal/repository/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	GetAllUsers() ([]models.User, error)
	SearchUsers(query string) ([]models.User, error)
}