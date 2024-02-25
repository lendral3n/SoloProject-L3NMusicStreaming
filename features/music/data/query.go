package data

import (
	"l3nmusic/features/music"

	"gorm.io/gorm"
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
func (repo *musicQuery) Insert(userIdLogin int, input music.Core) error {
	musicInput := CoreToModel(input)
	musicInput.UserID = uint(userIdLogin)

	tx := repo.db.Create(&musicInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAll implements music.MusicDataInterface.
func (repo *musicQuery) SelectAll() ([]music.Core, error) {
	panic("unimplemented")
}
