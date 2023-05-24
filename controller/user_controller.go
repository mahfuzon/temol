package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/temol/helper"
	"github.com/mahfuzon/temol/request/user_request"
	"github.com/mahfuzon/temol/response/api_response"
	"github.com/mahfuzon/temol/response/user_response"
	"github.com/mahfuzon/temol/service"
)

type UserController struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (userController *UserController) Register(ctx echo.Context) error {
	request := user_request.UserRegisterRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		fmt.Println("error binding")
		errorResponse := api_response.ConverseToErrorResponse("failed register", err.Error())
		return ctx.JSON(500, errorResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		fmt.Println("error validation")
		errorValidation := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		errorResponse := api_response.ConverseToErrorResponse("failed register", errorValidation)
		return ctx.JSON(422, errorResponse)
	}

	user, err := userController.userService.FindByEmailOrPhoneNumber(request.Email, request.PhoneNumber)
	if err != nil {
		fmt.Println("error check email")
		errorResponse := api_response.ConverseToErrorResponse("failed register", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	if user.Id > 0 {
		errorResponse := api_response.ConverseToErrorResponse("failed register", "user with this email or phone number already exists")
		return ctx.JSON(400, errorResponse)
	}

	passWordHash, err := userController.userService.HashPassword(request.Password)
	if err != nil {
		fmt.Println("error hash password")
		errorResponse := api_response.ConverseToErrorResponse("failed register", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	request.Password = passWordHash

	userProfileResponse, err := userController.userService.Register(request)
	if err != nil {
		fmt.Println("error register")
		errorResponse := api_response.ConverseToErrorResponse("failed register", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	successResponse := api_response.ConverseToSuccessResponse("success register", userProfileResponse)
	return ctx.JSON(200, successResponse)
}

func (userController *UserController) Login(ctx echo.Context) error {
	request := user_request.UserLoginRequest{}

	err := ctx.Bind(&request)
	if err != nil {
		fmt.Println("error binding")
		errorResponse := api_response.ConverseToErrorResponse("failed login", err.Error())
		return ctx.JSON(500, errorResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		fmt.Println("error validation")
		errorValidation := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		errorResponse := api_response.ConverseToErrorResponse("failed login", errorValidation)
		return ctx.JSON(422, errorResponse)
	}

	user, err := userController.userService.FindByEmailOrPhoneNumber(request.PhoneNumber, request.PhoneNumber)
	if err != nil {
		fmt.Println("error check email or phone number")
		errorResponse := api_response.ConverseToErrorResponse("failed login", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	if user.Id == 0 {
		errorResponse := api_response.ConverseToErrorResponse("failed login", "user belum terdaftar")
		return ctx.JSON(400, errorResponse)
	}

	ok, err := userController.userService.ValidatePassword(user, request.Password)
	if err != nil || !ok {
		errorResponse := api_response.ConverseToErrorResponse("failed login", "email or password invalid")
		return ctx.JSON(400, errorResponse)
	}

	token, err := userController.authService.GenerateAccessToken(user.Id)
	if err != nil {
		fmt.Println("error generate token")
		errorResponse := api_response.ConverseToErrorResponse("failed login", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	userLoginResponse := user_response.UserLoginResponse{
		Token: token,
	}

	successResponse := api_response.ConverseToSuccessResponse("success login", userLoginResponse)
	return ctx.JSON(200, successResponse)
}
