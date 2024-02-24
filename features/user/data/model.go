package data

import (
	"l3nmusic/features/user"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	UserName     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string `gorm:"not null"`
	Gender       string `gorm:"not null"`
	Role         string `gorm:"not null"`
	PhotoProfile string
}

func CoreToModel(input user.Core) User {
	return User{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		Password:     input.Password,
		Gender:       input.Gender,
		Role:         input.Role,
		PhotoProfile: input.PhotoProfile,
	}
}

func CoreToModelUpdate(input user.CoreUpdate) User {
	return User{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		PhotoProfile: input.PhotoProfile,
	}
}

func (u User) ModelToCore() user.Core {
	return user.Core{
		ID:           u.ID,
		Name:         u.Name,
		UserName:     u.UserName,
		Email:        u.Email,
		Password:     u.Password,
		Gender:       u.Gender,
		Role:         u.Role,
		PhotoProfile: u.PhotoProfile,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
