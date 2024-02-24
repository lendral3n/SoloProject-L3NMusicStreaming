package service

import (
	"l3nmusic/features/user"
	"l3nmusic/utils/encrypts"
	"l3nmusic/utils/middlewares"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("error hash password")
		}
		input.Password = hashedPass
	}

	err := service.userData.Insert(input)
	return err
}

// GetById implements user.UserServiceInterface.
func (service *userService) GetById(userId int) (*user.Core, error) {
	result, err := service.userData.SelectById(userId)
	return result, err
}

// Update implements user.UserServiceInterface.
func (service *userService) Update(userId int, input user.CoreUpdate) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	if userId <= 0 {
		return errors.New("invalid id.")
	}

	err := service.userData.Update(userId, input)
	return err
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(userId int) error {
	if userId <= 0 {
		return errors.New("invalid id")
	}
	err := service.userData.Delete(userId)
	return err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	if email == "" && password == "" {
		return nil, "", errors.New("email dan password wajib diisi.")
	}
	if email == "" {
		return nil, "", errors.New("email wajib diisi.")
	}
	if password == "" {
		return nil, "", errors.New("password wajib diisi.")
	}

	data, err = service.userData.Login(email, password)
	if err != nil {
		return nil, "", err
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai.")
	}

	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}
	return data, token, err
}

// ChangePassword implements user.UserServiceInterface.
func (service *userService) ChangePassword(userId int, oldPassword, newPassword string) error {
	user, errGet := service.GetById(userId)
	if errGet != nil {
		return errGet
	}

	if oldPassword == "" {
		return errors.New("please input current password")
	}

	if newPassword == "" {
		return errors.New("please input new password")
	}

	checkPassword := service.hashService.CheckPasswordHash(user.Password, oldPassword)
	if !checkPassword {
		return errors.New("current password not match")
	}

	checkNewPassword := service.hashService.CheckPasswordHash(user.Password, newPassword)
	if checkNewPassword {
		return errors.New("password cannot be the same")
	}

	hashedNewPass, errHash := service.hashService.HashPassword(newPassword)
	if errHash != nil {
		return errors.New("error hash password")
	}

	err := service.userData.ChangePassword(userId, oldPassword, hashedNewPass)
	return err

}
