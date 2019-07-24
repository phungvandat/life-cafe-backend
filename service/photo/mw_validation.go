package photo

import (
	"context"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
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

func (mw validationMiddleware) RollbackTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.RollbackTransaction(ctx, transactionID)
}

func (mw validationMiddleware) CommitTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.CommitTransaction(ctx, transactionID)
}

func (mw validationMiddleware) CreatePhoto(ctx context.Context, req requestModel.CreatePhotoRequest) (*responseModel.CreatePhotoResponse, error) {
	return mw.Service.CreatePhoto(ctx, req)
}

func (mw validationMiddleware) GetPhotos(ctx context.Context, req requestModel.GetPhotosRequest) (*responseModel.GetPhotosResponse, error) {
	return mw.Service.GetPhotos(ctx, req)
}

func (mw validationMiddleware) RemovePhoto(ctx context.Context, req requestModel.RemovePhotoRequest) (*responseModel.RemovePhotoResponse, error) {
	return mw.Service.RemovePhoto(ctx, req)
}

func (mw validationMiddleware) GetPhoto(ctx context.Context, req requestModel.GetPhotoRequest) (*responseModel.GetPhotoResponse, error) {
	return mw.Service.GetPhoto(ctx, req)
}
