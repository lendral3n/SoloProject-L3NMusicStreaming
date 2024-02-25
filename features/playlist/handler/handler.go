package handler

import (
	"l3nmusic/features/playlist"
	"l3nmusic/utils/middlewares"
	"l3nmusic/utils/responses"
	"net/http"

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
