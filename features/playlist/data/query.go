package data

import (
	"errors"
	"l3nmusic/app/cache"
	md "l3nmusic/features/music/data"
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
	playlistInput := CoreToModelPlaylist(input)
	playlistInput.UserID = uint(userIdLogin)

	tx := repo.db.Create(&playlistInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectPlaylist implements playlist.PlaylistDataInterface.
func (repo *playlistQuery) SelectPlaylistById(userIdLogin, playlistID int) (playlist.Core, error) {
	var playlist md.Playlist
	tx := repo.db.First(&playlist, playlistID)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return ModelToCorePlaylist(playlist), errors.New("playlist not found")
		}
		return ModelToCorePlaylist(playlist), tx.Error
	}

	if playlist.UserID != uint(userIdLogin) {
		return ModelToCorePlaylist(playlist), errors.New("playlist ini bukan milik anda")
	}

	return ModelToCorePlaylist(playlist), nil
}

// InsertSongToPlaylist implements playlist.PlaylistDataInterface.
func (repo *playlistQuery) InsertSongToPlaylist(input playlist.PlaylistSongCore) error {
	playlistInput := CoreToModelPlaylistSong(input)

	tx := repo.db.Create(&playlistInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
