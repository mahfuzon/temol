package user_request

type UserRegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"emailValidation"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number" validate:"required,number,min=11"`
	Address     string `json:"address" validate:"required"`
}
