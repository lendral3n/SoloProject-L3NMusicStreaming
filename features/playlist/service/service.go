package service

import (
	"l3nmusic/features/playlist"
)

type playlistService struct {
	musicData playlist.PlaylistDataInterface
}

// dependency injection
func New(repo playlist.PlaylistDataInterface) playlist.PlaylistServiceInterface {
	return &playlistService{
		musicData: repo,
	}
}

// Create implements playlist.PlaylistServiceInterface.
func (service *playlistService) Create(userIdLogin int, input playlist.Core) error {
	err := service.musicData.Insert(userIdLogin, input)
	if err != nil {
		return err
	}

	return nil
}
