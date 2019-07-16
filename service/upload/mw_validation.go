package upload

import (
	"context"
	"net/http"
	"strings"

	serviceModel "github.com/phungvandat/life-cafe-backend/model/service"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware for validation purposes
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) UploadImages(ctx context.Context, inputFiles []serviceModel.UploadFileInput) ([]string, error) {
	for _, image := range inputFiles {
		defer (*image.File).Close()
		buffer := make([]byte, 512)
		_, err := (*image.File).Read(buffer)
		if err != nil {
			return []string{}, err
		}
		(*image.File).Seek(0, 0)
		contentType := strings.Split(http.DetectContentType(buffer), "/")
		if contentType[0] != "image" {
			return []string{}, FileAreNotImageTypes
		}
	}
	return mw.Service.UploadImages(ctx, inputFiles)
}
