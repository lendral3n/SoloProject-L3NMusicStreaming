package handler

import "l3nmusic/features/music"

type MusicRequest struct {
	UserID     uint
	Title      string  `json:"title" form:"title"`
	Artist     string  `json:"artis" form:"artis"`
	Genre      string  `json:"genre" form:"genre"`
	Duration   float64 `json:"duration" form:"duration"`
	Music      string  `json:"music" form:"music"`
	PhotoMusic string  `json:"photo_music" form:"photo_music"`
}

func RequestToCore(input MusicRequest, musicURL, imageURL string, userIdLogin uint) music.Core {
	return music.Core{
		UserID:     userIdLogin,
		Title:      input.Title,
		Artist:     input.Artist,
		Genre:      input.Genre,
		Duration:   input.Duration,
		Music:      musicURL,
		PhotoMusic: imageURL,
	}
}
