package photo

import (
	"context"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
)

// Service interface contain functions
type Service interface {
	RollbackTransaction(context.Context, string) error
	CommitTransaction(context.Context, string) error
	CreatePhoto(context.Context, requestModel.CreatePhotoRequest) (*responseModel.CreatePhotoResponse, error)
	GetPhotos(context.Context, requestModel.GetPhotosRequest) (*responseModel.GetPhotosResponse, error)
	RemovePhoto(context.Context, requestModel.RemovePhotoRequest) (*responseModel.RemovePhotoResponse, error)
	GetPhoto(context.Context, requestModel.GetPhotoRequest) (*responseModel.GetPhotoResponse, error)
}
