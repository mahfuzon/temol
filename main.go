package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mahfuzon/temol/controller"
	"github.com/mahfuzon/temol/libraries"
	"github.com/mahfuzon/temol/repository"
	"github.com/mahfuzon/temol/service"
	"github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db := libraries.SetDb()
	log := libraries.NewLogger()
	userRepository := repository.NewUserRepository(db, log)
	userService := service.NewUserService(userRepository, log)
	authService := service.NewAuthService()
	userController := controller.NewUserController(userService, authService, log)
	router := libraries.SetRouter()

	router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))

	api := router.Group("/api")

	apiV1 := api.Group("/v1")

	apiV1Auth := apiV1.Group("/auth")
	apiV1Auth.POST("/register", userController.Register)
	apiV1Auth.POST("/login", userController.Login)

	router.Logger.Fatal(router.Start(":8000"))
}
