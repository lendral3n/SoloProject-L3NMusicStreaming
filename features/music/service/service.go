package service

import (
	"context"
	"errors"
	"l3nmusic/features/music"
	"l3nmusic/features/user"
)

type musicService struct {
	musicData   music.MusicDataInterface
	userService user.UserServiceInterface
}

// dependency injection
func New(repo music.MusicDataInterface, us user.UserServiceInterface) music.MusicServiceInterface {
	return &musicService{
		musicData:   repo,
		userService: us,
	}
}

// Create implements music.MusicServiceInterface.
func (service *musicService) Create(ctx context.Context, userIdLogin int, input music.Core) error {
	user, err := service.userService.GetById(ctx, userIdLogin)
	if err != nil {
		return err
	}

	if user.Role != "admin" {
		return errors.New("anda tidak memiliki akses untuk fitur ini")
	}

	err = service.musicData.Insert(userIdLogin, input)
	if err != nil {
		return err
	}

	return nil
}

// SelectAll implements music.MusicServiceInterface.
func (service *musicService) GetAll(ctx context.Context, page int, limit int) ([]music.Core, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	result, err := service.musicData.SelectAll(ctx, page, limit)
	return result, err
}

// AddLikedSong implements music.MusicServiceInterface.
func (service *musicService) AddLikedSong(userIdLogin int, songId int) (string, error) {
	isLiked, err := service.musicData.CheckLikedSong(userIdLogin, songId)
	if err != nil {
		return "", err
	}
	if isLiked {
		err = service.musicData.DeleteLikedSong(userIdLogin, songId)
		if err != nil {
			return "", err
		}
		return "berhasil menghapus musik yang disukai", nil
	} else {
		err = service.musicData.InsertLikedSong(userIdLogin, songId)
		if err != nil {
			return "", err
		}
		return "berhasil menambahkan musik yang disukai", nil
	}
}
