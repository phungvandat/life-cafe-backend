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
		var inputFiles = request.(requestModel.UploadImagesRequest)
		links, err := s.UploadService.UploadImages(ctx, inputFiles)
		if err != nil {
			return nil, err
		}

		return responseModel.UploadImagesResponse{
			Links: links,
		}, nil
	}
}

// MakeGetImageFile func
func MakeGetImageFile(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var path = request.(string)
		file, err := s.UploadService.GetImageFile(ctx, path)

		if err != nil {
			return nil, err
		}

		return file, nil
	}
}
