package data

import (
	"l3nmusic/app/cache"
	"l3nmusic/features/playlist"

	"gorm.io/gorm"
)

type playlistQuery struct {
	db    *gorm.DB
	redis cache.Redis
}

func New(db *gorm.DB, redis cache.Redis) playlist.PlaylistDataInterface {
	return &playlistQuery{
		db:    db,
		redis: redis,
	}
}

// Insert implements playlist.PlaylistDataInterface.
func (repo *playlistQuery) Insert(userIdLogin int, input playlist.Core) error {
	playlistInput := CoreToModel(input)
	playlistInput.UserID = uint(userIdLogin)

	tx := repo.db.Create(&playlistInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}