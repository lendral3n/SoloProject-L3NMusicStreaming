package cloudinary

import (
	"context"
	"fmt"
	"l3nmusic/app/config"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryUploaderInterface interface {
	UploadImage(fileHeader *multipart.FileHeader) (string, error)
}

type CloudinaryUploader struct {
}

func New() CloudinaryUploaderInterface {
	return &CloudinaryUploader{}
}

func (cu *CloudinaryUploader) UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	cld, err := cloudinary.NewFromURL(config.CLD_URL)
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		return "", fmt.Errorf("invalid file type: %w", err)
	}

	uploadParams := uploader.UploadParams{
		Folder: "L3N_Music",
	}

	resp, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", fmt.Errorf("error uploading to Cloudinary: %w", err)
	}

	return resp.SecureURL, nil
}
