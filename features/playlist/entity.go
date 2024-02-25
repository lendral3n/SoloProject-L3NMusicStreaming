package playlist

import "time"

type Core struct {
	ID        uint
	Name      string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// interface untuk Data Layer
type PlaylistDataInterface interface {
	Insert(userIdLogin int, input Core) error
}

// interface untuk Service Layer
type PlaylistServiceInterface interface {
	Create(userIdLogin int, input Core) error
}
