package main

import (
	"github.com/joho/godotenv"
	"github.com/mahfuzon/temol/controller"
	"github.com/mahfuzon/temol/libraries"
	"github.com/mahfuzon/temol/repository"
	"github.com/mahfuzon/temol/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db := libraries.SetDb()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	router := libraries.SetRouter()

	api := router.Group("/api")

	apiV1 := api.Group("/v1")

	apiV1Auth := apiV1.Group("/auth")
	apiV1Auth.POST("/register", userController.Register)
	apiV1Auth.POST("/login", userController.Login)
}
