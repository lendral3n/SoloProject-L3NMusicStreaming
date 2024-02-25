package handler

import "l3nmusic/features/playlist"

type PlaylistResponse struct {
	ID   uint   `json:"playlist_id"`
	Name string `json:"name_playlist"`
}

func CoreToResponsePlaylist(data playlist.Core) PlaylistResponse {
	return PlaylistResponse{
		ID:   data.ID,
		Name: data.Name,
	}
}