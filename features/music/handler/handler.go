package handler

import (
	"l3nmusic/features/music"
)

type MusicHandler struct {
	musicService music.MusicServiceInterface
}

func New(service music.MusicServiceInterface) *MusicHandler {
	return &MusicHandler{
		musicService: service,
	}
}
