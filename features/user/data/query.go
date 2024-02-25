package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"l3nmusic/app/cache"
	"l3nmusic/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db    *gorm.DB
	redis cache.Redis
}

func New(db *gorm.DB, redis cache.Redis) user.UserDataInterface {
	return &userQuery{
		db:    db,
		redis: redis,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	dataGorm := CoreToModel(input)

	tx := repo.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectById implements user.UserDataInterface.
func (repo *userQuery) SelectById(ctx context.Context, userId int) (*user.Core, error) {
	userData, err := repo.redis.Get(ctx, fmt.Sprintf("user:%d", userId))
	if err == nil && userData != "" {
		var user user.Core
		err = json.Unmarshal([]byte(userData), &user)
		if err == nil {
			return &user, nil
		}
	}

	var userDataGorm User
	tx := repo.db.First(&userDataGorm, userId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := userDataGorm.ModelToCore()

	jsonData, err := json.Marshal(result)
	if err == nil {
		userData = string(jsonData)
		err = repo.redis.Set(ctx, fmt.Sprintf("user:%d", userId), userData)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(userId int, input user.CoreUpdate) error {
	dataGorm := CoreToModelUpdate(input)
	tx := repo.db.Model(&User{}).Where("id = ?", userId).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(userId int) error {
	tx := repo.db.Delete(&User{}, userId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (data *user.Core, err error) {
	var userGorm User
	tx := repo.db.Where("email = ?", email).First(&userGorm)
	if tx.Error != nil {
		// return nil, tx.Error
		return nil, errors.New(" Invalid email or password")
	}
	result := userGorm.ModelToCore()
	return &result, nil
}

// ChangePassword implements user.UserDataInterface.
func (repo *userQuery) ChangePassword(userId int, oldPassword, newPassword string) error {
	var userGorm User
	userGorm.Password = newPassword
	tx := repo.db.Model(&User{}).Where("id = ?", userId).Updates(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

func (repo *userQuery) GetTotalUser() (int, error) {
	var count int64
	tx := repo.db.Model(&User{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}
