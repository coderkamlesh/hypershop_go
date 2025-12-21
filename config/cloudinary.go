package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

var CloudinaryInstance *cloudinary.Cloudinary

func SetupCloudinary() error {
	cld, err := cloudinary.NewFromParams(
		AppConfig.CloudinaryCloudName,
		AppConfig.CloudinaryAPIKey,
		AppConfig.CloudinaryAPISecret,
	)
	if err != nil {
		return err
	}

	CloudinaryInstance = cld
	return nil
}
