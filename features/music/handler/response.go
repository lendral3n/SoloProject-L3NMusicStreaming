package handler

import (
	"l3nmusic/features/music"
)

type MusicResponse struct {
	ID uint `json:"id" form:"id"`
	Title      string  `json:"title" form:"title"`
	Artist     string  `json:"artis" form:"artis"`
	Genre      string  `json:"genre" form:"genre"`
	Duration   float64 `json:"duration" form:"duration"`
	Music      string  `json:"music" form:"music"`
	PhotoMusic string  `json:"photo_music" form:"photo_music"`
}

func CoreToResponseMusic(data music.Core) MusicResponse {
	return MusicResponse{
		ID:         data.ID,
		Title:      data.Title,
		Artist:     data.Artist,
		Genre:      data.Genre,
		Duration:   data.Duration,
		Music:      data.Music,
		PhotoMusic: data.PhotoMusic,
	}
}
