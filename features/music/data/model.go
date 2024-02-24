package data

import (
	"l3nmusic/features/music"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title    string
	Artist   string
	Genre    string
	Duration int
}

func CoreToModel(input music.Core) Song {
	return Song{
		Title:    input.Title,
		Artist:   input.Artist,
		Genre:    input.Genre,
		Duration: input.Duration,
	}
}

func (s Song) ModelToCore() music.Core {
	return music.Core{
		ID:        s.ID,
		Title:     s.Title,
		Artist:    s.Artist,
		Genre:     s.Genre,
		Duration:  s.Duration,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}
