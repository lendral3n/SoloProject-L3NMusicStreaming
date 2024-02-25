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
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	User      user.Core
}

// interface untuk Data Layer
type MusicDataInterface interface {
	Insert(userIdLogin int, input Core) error
	SelectAll() ([]Core, error)
}

// interface untuk Service Layer
type MusicServiceInterface interface {
	Create(ctx context.Context, userIdLogin int, input Core) error
	SelectAll() ([]Core, error)
}
