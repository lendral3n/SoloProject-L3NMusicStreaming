package handler

import "l3nmusic/features/playlist"

type PlaylistRequest struct {
	UserID uint
	Name   string `json:"name_playlist" form:"name_playlist"`
}

type AddSongRequest struct {
	SongID uint `json:"song_id"`
}

func RequestToCore(input PlaylistRequest, userIdLogin uint) playlist.Core {
	return playlist.Core{
		UserID: userIdLogin,
		Name:   input.Name,
	}
}