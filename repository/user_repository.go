package repository

import (
	"github.com/mahfuzon/temol/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByEmailOrPhoneNumber(email, phoneNumber string) (models.User, error)
	FindById(id int) (models.User, error)
}

type userRepository struct {
	log *logrus.Logger
	db  *gorm.DB
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) *userRepository {
	return &userRepository{
		log: log,
		db:  db,
	}
}

func (repository *userRepository) Create(user models.User) (models.User, error) {
	repository.log.Info("masuk ke user repository create")
	err := repository.db.Create(&user).Error
	if err != nil {
		repository.log.Error(err.Error())
		return user, err
	}

	repository.log.Info("success userRepository.Create")
	return user, nil
}

func (repository *userRepository) FindByEmailOrPhoneNumber(email, phoneNumber string) (models.User, error) {
	repository.log.Info("userRepository.FindByEmailOrPhoneNumber")
	user := models.User{}
	err := repository.db.Where("email = ?", email).Or("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		repository.log.Error(err.Error())
		return user, err
	}

	repository.log.Info("success userRepository.FindByEmailOrPhoneNumber")
	return user, nil
}

func (repository *userRepository) FindById(id int) (models.User, error) {
	repository.log.Info("masuk ke user repository FindById")
	user := models.User{}
	err := repository.db.First(&user, id).Error
	if err != nil {
		repository.log.Info(err.Error())
		return user, err
	}

	repository.log.Info("success find by id")
	return user, nil
}
