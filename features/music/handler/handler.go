package handler

import (
	"l3nmusic/features/music"
	"l3nmusic/utils/middlewares"
	"l3nmusic/utils/responses"
	"l3nmusic/utils/upload"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MusicHandler struct {
	musicService music.MusicServiceInterface
	s3           upload.S3UploaderInterface
}

func New(service music.MusicServiceInterface, s3Uploader upload.S3UploaderInterface) *MusicHandler {
	return &MusicHandler{
		musicService: service,
		s3:           s3Uploader,
	}
}

func (handler *MusicHandler) CreateMusic(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}
	
	ctx := c.Request().Context()
	newMusic := MusicRequest{}
	errBind := c.Bind(&newMusic)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid "+errBind.Error(), nil))
	}

	fileHeader, err := c.FormFile("music")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error reading music file: "+err.Error(), nil))
	}

	musicURL, err := handler.s3.UploadMusic(fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading music to S3: "+err.Error(), nil))
	}

	fileData, err := c.FormFile("photo_music")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the file", nil))
	}

	var imageURL string
	if fileData != nil {
		imageURL, err = handler.s3.UploadImage(fileData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading the image "+err.Error(), nil))
		}
	}

	musicCore := RequestToCore(newMusic, musicURL, imageURL, uint(userIdLogin))

	errInsert := handler.musicService.Create(ctx, userIdLogin, musicCore)
	if errInsert != nil {
		if errInsert.Error() == "anda tidak memiliki akses untuk fitur ini" {
			return c.JSON(http.StatusUnauthorized, responses.WebResponse(errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("berhasil menambahkan musik", nil))
}
