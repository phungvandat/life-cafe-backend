package upload

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/service"
)

// MakeUploadImages func
func MakeUploadImages(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			inputFiles = request.(requestModel.UploadImagesRequest)
		)
		link, err := s.UploadService.UploadImages(ctx, inputFiles)
		if err != nil {
			return nil, err
		}

		return responseModel.UploadImagesResponse{
			Link: link,
		}, nil
	}
}
