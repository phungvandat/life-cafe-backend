package upload

import (
	"context"

	serviceModel "github.com/phungvandat/life-cafe-backend/model/service"
)

// Service interface contain functions
type Service interface {
	UploadImages(ctx context.Context, inputFiles []serviceModel.UploadFileInput) ([]string, error)
}
