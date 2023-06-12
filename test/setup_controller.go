package test

import (
	"github.com/mahfuzon/temol/controller"
	"github.com/mahfuzon/temol/repository"
	"github.com/mahfuzon/temol/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupUserController(db *gorm.DB, log *logrus.Logger) *controller.UserController {
	userRepository := repository.NewUserRepository(db, log)
	userService := service.NewUserService(userRepository, log)
	authService := service.NewAuthService()
	userController := controller.NewUserController(userService, authService, log)
	return userController
}
