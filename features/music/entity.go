package music

import (
	"context"
	"l3nmusic/features/user"
	"time"
)

type Core struct {
	ID         uint
	Title      string
	Artist     string
	Genre      string
	Duration   float64
	Music      string
	PhotoMusic string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     uint
	User       user.Core
	Likes      []CoreLiked
}

type CoreLiked struct {
	SongID uint
	UserID uint
    CreatedAt  time.Time
	UpdatedAt  time.Time
}

// interface untuk Data Layer
type MusicDataInterface interface {
	Insert(userIdLogin int, input Core) error
	SelectAll(ctx context.Context, page, limit int) ([]Core, error)
	InsertLikedSong(userIdLogin, songId int) error
	CheckLikedSong(userIdLogin, songId int) (bool, error)
	DeleteLikedSong(userIdLogin, songId int) error
	SelectLikedSong(ctx context.Context, userIdLogin, page, limit int) ([]Core, error)
	SearchMusic(ctx context.Context, query string, page, limit int) ([]Core, error)
}

// interface untuk Service Layer
type MusicServiceInterface interface {
	Create(ctx context.Context, userIdLogin int, input Core) error
	GetAll(ctx context.Context, page, limit int) ([]Core, error)
	AddLikedSong(userIdLogin, songId int) (string, error)
	GetLikedSong(ctx context.Context, userIdLogin, page, limit int) ([]Core, error)
	SearchMusic(ctx context.Context, query string, page, limit int) ([]Core, error)
}
