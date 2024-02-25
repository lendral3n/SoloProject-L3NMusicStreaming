package data

import (
	md "l3nmusic/features/music/data"
	"l3nmusic/features/playlist"

	"gorm.io/gorm"
)

type PlaylistSong struct {
	gorm.Model
	PlaylistID uint
	SongID     uint
}

func CoreToModelPlaylist(input playlist.Core) md.Playlist {
	return md.Playlist{
		Name:   input.Name,
		UserID: input.UserID,
	}
}

func CoreToModelPlaylistSong(input playlist.PlaylistSongCore)PlaylistSong {
	return PlaylistSong{
		PlaylistID:   input.PlaylistID,
		SongID: input.SongID,
	}
}

func ModelToCorePlaylist(p md.Playlist) playlist.Core {
	return playlist.Core{
		ID:        p.ID,
		Name:      p.Name,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
