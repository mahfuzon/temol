package user_response

import "github.com/mahfuzon/temol/models"

type UserProfileResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	ProfileImage string `json:"profile_image"`
	Address      string `json:"address"`
}

func ConverseToUserProfileResponse(user models.User) UserProfileResponse {
	userProfileResponse := UserProfileResponse{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email.String,
		PhoneNumber:  user.PhoneNumber.String,
		ProfileImage: user.ProfileImage,
		Address:      user.Address,
	}

	return userProfileResponse
}
