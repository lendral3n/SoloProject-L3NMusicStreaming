package playlist

import (
	"context"
	"l3nmusic/features/music"
	"time"
)

type Core struct {
	ID        uint
	Name      string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PlaylistSongCore struct {
	ID        uint
	PlaylistID uint
	SongID     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// interface untuk Data Layer
type PlaylistDataInterface interface {
	Insert(userIdLogin int, input Core) error
	InsertSongToPlaylist(input PlaylistSongCore) error
	SelectPlaylistById(userIdLogin, playlistID int) (Core, error)
	SelectPlaylistsByUser(userIdLogin int) ([]Core, error)
	SelectSongsInPlaylist(ctx context.Context, playlistID int) ([]music.Core, error)
	DeleteSongFromPlaylist(playlistID, songID int) error
	DeletePlaylist(userIdLogin, playlistID int) error
}

// interface untuk Service Layer
type PlaylistServiceInterface interface {
	Create(userIdLogin int, input Core) error
	CreateSongToPlaylist(userIdLogin int, input PlaylistSongCore) error
	GetUserPlaylists(userIdLogin int) ([]Core, error)
	GetSongsInPlaylist(ctx context.Context, playlistID int) ([]music.Core, error)
	DeletePlaylist(userIdLogin, playlistID int) error
	DeleteSongFromPlaylist(userIdLogin, playlistID, songID int) error
}
