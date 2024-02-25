package data

import (
	"l3nmusic/features/music"
	"l3nmusic/features/user/data"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title      string
	Artist     string
	Genre      string
	Duration   float64
	Music      string
	PhotoMusic string
	UserID     uint
	User       data.User
	Likes      []LikedSong
}

type LikedSong struct {
	gorm.Model
	UserID uint
	User   data.User
	SongID uint
	Song   Song
}

func CoreToModel(input music.Core) Song {
	return Song{
		UserID:     input.UserID,
		Title:      input.Title,
		Artist:     input.Artist,
		Genre:      input.Genre,
		Duration:   input.Duration,
		Music:      input.Music,
		PhotoMusic: input.PhotoMusic,
	}
}

func (s Song) ModelToCore() music.Core {
	return music.Core{
		ID:         s.ID,
		Title:      s.Title,
		Artist:     s.Artist,
		Genre:      s.Genre,
		Duration:   s.Duration,
		Music:      s.Music,
		PhotoMusic: s.PhotoMusic,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
	}
}

func CoreToModelLiked(input music.CoreLiked) LikedSong {
	return LikedSong{
		UserID: input.UserID,
		SongID: input.SongID,
	}
}
