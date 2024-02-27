package data

import (
	"context"
	"errors"
	"l3nmusic/app/cache"
	"l3nmusic/features/music"
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

// SelectPlaylistsByUser implements playlist PlaylistDataInterface.
func (repo *playlistQuery) SelectPlaylistsByUser(userIdLogin int) ([]playlist.Core, error) {
	var playlists []md.Playlist
	tx := repo.db.Where("user_id = ?", userIdLogin).Find(&playlists)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []playlist.Core
	for _, p := range playlists {
		result = append(result, ModelToCorePlaylist(p))
	}

	return result, nil
}

// SelectSongsInPlaylist implements playlist.PlaylistDataInterface.
func (repo *playlistQuery) SelectSongsInPlaylist(ctx context.Context, playlistID int) ([]music.Core, error) {
	var playlistSongs []PlaylistSong
	tx := repo.db.Where("playlist_id = ?", playlistID).Find(&playlistSongs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var songs []music.Core
	for _, playlistSong := range playlistSongs {
		var song md.Song
		tx := repo.db.First(&song, playlistSong.SongID)
		if tx.Error != nil {
			return nil, tx.Error
		}
		songs = append(songs, song.ModelToCore())
	}

	return songs, nil
}

// DeletePlaylist implements playlist.PlaylistDataInterface.
func (repo *playlistQuery) DeletePlaylist(userIdLogin int, playlistID int) error {
	var playlist md.Playlist
	tx := repo.db.First(&playlist, playlistID)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return errors.New("playlist not found")
		}
		return tx.Error
	}

	// Delete songs from the playlist
	tx = repo.db.Where("playlist_id = ?", playlistID).Delete(&PlaylistSong{})
	if tx.Error != nil {
		return tx.Error
	}

	// Delete the playlist
	tx = repo.db.Where("id = ?", playlistID).Delete(&md.Playlist{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteSongFromPlaylist implements playlist.PlaylistDataInterface.
func (repo *playlistQuery) DeleteSongFromPlaylist(playlistID int, songID int) error {
	tx := repo.db.Where("playlist_id = ? AND song_id = ?", playlistID, songID).Delete(&PlaylistSong{})
    if tx.Error != nil {
        return tx.Error
    }
    return nil
}