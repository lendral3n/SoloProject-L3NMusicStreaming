package handler

import "l3nmusic/features/user"

type UserRequest struct {
	Name         string `json:"name" form:"name"`
	UserName     string `json:"user_name" form:"user_name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Gender       string `json:"gender" form:"gender"`
	Role         string `json:"role" form:"role"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		Password:     input.Password,
		Gender:       input.Gender,
		Role:         input.Role,
		PhotoProfile: input.PhotoProfile,
	}
}

func UpdateRequestToCore(input UserRequest, imageURL string) user.Core {
	return user.Core{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		Gender:       input.Gender,
		PhotoProfile: imageURL,
	}
}

func UpdateRequestToCoreUpdate(input UserRequest, imageURL string) user.CoreUpdate {
	return user.CoreUpdate{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		PhotoProfile: imageURL,
	}
}
