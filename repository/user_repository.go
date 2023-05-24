package repository

import (
	"github.com/mahfuzon/temol/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByEmailOrPhoneNumber(email, phoneNumber string) (models.User, error)
	FindById(id int) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repository *userRepository) Create(user models.User) (models.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepository) FindByEmailOrPhoneNumber(email, phoneNumber string) (models.User, error) {
	user := models.User{}
	err := repository.db.Where("email = ?", email).Or("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepository) FindById(id int) (models.User, error) {
	user := models.User{}
	err := repository.db.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
