package repositories

import (
	"errors"
	"qc/internal/repository/models"
	"qc/internal/repository/rInterfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// type UserRepository interface {
// 	CreateUser(user *models.User) error
// 	DeleteUser(id int) error
// 	GetUserByUsername(username string) (*models.User, error)
// 	GetUserByEmail(email string) (*models.User, error)
// 	UpdateUser(user *models.User) error
// 	GetAllUsers() ([]models.User, error)
// 	SearchUsers(query string) ([]models.User, error)
// }

func NewUserRepository(db *gorm.DB) rInterfaces.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) DeleteUser(id int) error {
	result := r.db.Delete(&models.User{}, id)
	if result.RowsAffected == 0 {
		return errors.New("пользователь не найден")
	}
	return result.Error
	
}

