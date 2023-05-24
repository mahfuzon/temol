package test

import (
	"github.com/mahfuzon/temol/controller"
	"github.com/mahfuzon/temol/repository"
	"github.com/mahfuzon/temol/service"
	"gorm.io/gorm"
)

func SetupUserController(db *gorm.DB) *controller.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController
}
