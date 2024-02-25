package playlist

import "time"

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
}

// interface untuk Service Layer
type PlaylistServiceInterface interface {
	Create(userIdLogin int, input Core) error
	CreateSongToPlaylist(userIdLogin int, input PlaylistSongCore) error
}
