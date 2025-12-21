package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/coderkamlesh/hypershop_go/config"
)

func BoolPtr(b bool) *bool {
	return &b
}

// UploadFile uploads a file to Cloudinary with folder organization
func UploadFile(file multipart.File, filename string, folder string) (string, error) {
	ctx := context.Background()

	// If folder is empty, use default root folder
	if folder == "" {
		folder = "hypershop/general"
	}

	// Generate unique filename to avoid conflicts
	uniqueFilename := GenerateUniqueFilename(filename)

	uploadParams := uploader.UploadParams{
		PublicID:       uniqueFilename,
		Folder:         folder,
		ResourceType:   "auto",
		UniqueFilename: BoolPtr(false), // Use pointer to bool
		Overwrite:      BoolPtr(false),
	}

	result, err := config.CloudinaryInstance.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}

// GenerateUniqueFilename creates a unique filename with timestamp
func GenerateUniqueFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	nameWithoutExt := strings.TrimSuffix(originalFilename, ext)
	timestamp := time.Now().Unix()

	return fmt.Sprintf("%s_%d%s", nameWithoutExt, timestamp, ext)
}

// UploadMultipleFiles uploads multiple files to the same folder
func UploadMultipleFiles(files []*multipart.FileHeader, folder string) ([]string, error) {
	var urls []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}

		url, err := UploadFile(file, fileHeader.Filename, folder)
		file.Close()

		if err != nil {
			continue
		}

		urls = append(urls, url)
	}

	return urls, nil
}
