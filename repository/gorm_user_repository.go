package repository

import (
	"praktikum/dto"
	"praktikum/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(dataUser model.User) error
	LoginUser(userLog dto.UserDTO) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) CreateUser(dataUser model.User) error {
	if err := r.db.Create(&dataUser).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) LoginUser(userLog dto.UserDTO) (model.User, error) {
	var user model.User
	if err := r.db.Where("email = ? AND password = ?", userLog.Email, userLog.Password).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
