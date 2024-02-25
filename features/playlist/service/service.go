package service

import (
	"errors"
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

// CreateSongToPlaylist implements playlist.PlaylistServiceInterface.
func (service *playlistService) CreateSongToPlaylist(userIdLogin int, input playlist.PlaylistSongCore) error {
	playlist, err := service.musicData.SelectPlaylistById(userIdLogin, int(input.PlaylistID))
	if err != nil {
		return err
	}

	if playlist.UserID != uint(userIdLogin) {
		return errors.New("playlist ini bukan milik anda")
	}

	err = service.musicData.InsertSongToPlaylist(input)
	if err != nil {
		return err
	}
	return nil
}
