package service

import (
	"l3nmusic/features/music"
)

type musicService struct {
	musicData music.MusicDataInterface
}

// dependency injection
func New(repo music.MusicDataInterface) music.MusicServiceInterface {
	return &musicService{
		musicData: repo,
	}
}

// Create implements music.MusicServiceInterface.
func (service *musicService) Create(input music.Core) error {
	panic("unimplemented")
}

// SelectAll implements music.MusicServiceInterface.
func (service *musicService) SelectAll() ([]music.Core, error) {
	panic("unimplemented")
}
