package service

import (
	"errors"
	"l3nmusic/features/playlist"
)

type playlistService struct {
	playlistData playlist.PlaylistDataInterface
}

// dependency injection
func New(repo playlist.PlaylistDataInterface) playlist.PlaylistServiceInterface {
	return &playlistService{
		playlistData: repo,
	}
}

// Create implements playlist.PlaylistServiceInterface.
func (service *playlistService) Create(userIdLogin int, input playlist.Core) error {
	err := service.playlistData.Insert(userIdLogin, input)
	if err != nil {
		return err
	}

	return nil
}

// CreateSongToPlaylist implements playlist.PlaylistServiceInterface.
func (service *playlistService) CreateSongToPlaylist(userIdLogin int, input playlist.PlaylistSongCore) error {
	playlist, err := service.playlistData.SelectPlaylistById(userIdLogin, int(input.PlaylistID))
	if err != nil {
		return err
	}

	if playlist.UserID != uint(userIdLogin) {
		return errors.New("playlist ini bukan milik anda")
	}

	err = service.playlistData.InsertSongToPlaylist(input)
	if err != nil {
		return err
	}
	return nil
}

// GetUserPlaylists implements playlist.PlaylistServiceInterface.
func (service *playlistService) GetUserPlaylists(userIdLogin int) ([]playlist.Core, error) {
	result, err := service.playlistData.SelectPlaylistsByUser(userIdLogin)
	return result, err
}