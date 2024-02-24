package music

import "time"

type Core struct {
	ID        uint
	Title     string
	Artist    string
	Genre     string
	Duration  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// interface untuk Data Layer
type MusicDataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
}

// interface untuk Service Layer
type MusicServiceInterface interface {
	Create(input Core) error
	SelectAll() ([]Core, error)
}
