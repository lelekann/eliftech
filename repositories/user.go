package repositories

import (
	"eliftech_1/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdteUser(id uint, user *models.CreateUserInput) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	res := r.db.Find(&users)
	return users, res.Error
}

func (r *userRepository) GetUserById(id uint) (*models.User, error) {
	var user *models.User
	res := r.db.First(&user, id)
	return user, res.Error
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	res := r.db.Create(&user)
	return user, res.Error
}

func (r *userRepository) UpdteUser(id uint, user *models.CreateUserInput) error {
	res := r.db.Model(&user).Where("id=?", id).Update("user", user)
	return res.Error
}

func (r *userRepository) DeleteUser(id uint) error {
	res := r.db.Where(id).Delete(&models.User{})
	return res.Error
}
