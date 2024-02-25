package handler

import (
	"l3nmusic/features/playlist"
	"l3nmusic/utils/middlewares"
	"l3nmusic/utils/responses"
	"net/http"
	mh "l3nmusic/features/music/handler"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PlaylistHandler struct {
	playlistService playlist.PlaylistServiceInterface
}

func New(service playlist.PlaylistServiceInterface) *PlaylistHandler {
	return &PlaylistHandler{
		playlistService: service,
	}
}

func (handler *PlaylistHandler) CreatePlaylist(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newPlaylist := PlaylistRequest{}
	errBind := c.Bind(&newPlaylist)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid "+errBind.Error(), nil))
	}

	playlistCore := RequestToCore(newPlaylist, uint(userIdLogin))
	errInsert := handler.playlistService.Create(userIdLogin, playlistCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("berhasil membuat playlist", nil))
}

func (handler *PlaylistHandler) AddSongToPlaylist(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	playlistID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid playlist ID", nil))
	}

	var request AddSongRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Error binding data", nil))
	}

	input := playlist.PlaylistSongCore{
		PlaylistID: uint(playlistID),
		SongID:     request.SongID,
	}

	if err := handler.playlistService.CreateSongToPlaylist(userIdLogin, input); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil menambahkan lagu ke playlist", nil))
}

func (handler *PlaylistHandler) GetUserPlaylists(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	playlists, err := handler.playlistService.GetUserPlaylists(userIdLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var response []PlaylistResponse
	for _, p := range playlists {
		response = append(response, CoreToResponsePlaylist(p))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mendapatkan playlist", response))
}

func (handler *PlaylistHandler) GetSongsInPlaylist(c echo.Context) error {
	// userIdLogin := middlewares.ExtractTokenUserId(c)
	// if userIdLogin == 0 {
	// 	return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	// }

	ctx := c.Request().Context()
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid playlist ID", nil))
	}

	songs, err := handler.playlistService.GetSongsInPlaylist(ctx, playlistID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var response []mh.MusicResponse
	for _, s := range songs {
		response = append(response, mh.CoreToResponseMusic(s))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mendapatkan lagu dalam playlist", response))
}
