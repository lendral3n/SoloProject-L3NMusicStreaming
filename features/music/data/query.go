package data

import (
	"gorm.io/gorm"
	"l3nmusic/features/music"
)

type musicQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) music.MusicDataInterface {
	return &musicQuery{
		db: db,
	}
}


// Insert implements music.MusicDataInterface.
func (repo *musicQuery) Insert(input music.Core) error {
	panic("unimplemented")
}

// SelectAll implements music.MusicDataInterface.
func (repo *musicQuery) SelectAll() ([]music.Core, error) {
	panic("unimplemented")
}