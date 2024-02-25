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
func (service *musicService) SelectAll() ([]music.Core, error) {
	panic("unimplemented")
}
