package product

import (
	"context"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
)

// Service interface contain functions
type Service interface {
	RollbackTransaction(context.Context, string) error
	CommitTransaction(context.Context, string) error
	CreateProduct(context.Context, requestModel.CreateProductRequest) (*responseModel.CreateProductResponse, error)
	GetProduct(context.Context, requestModel.GetProductRequest) (*responseModel.GetProductResponse, error)
	GetProducts(context.Context, requestModel.GetProductsRequest) (*responseModel.GetProductsResponse, error)
	UpdateProduct(context.Context, requestModel.UpdateProductRequest) (*responseModel.UpdateProductResponse, error)
}
