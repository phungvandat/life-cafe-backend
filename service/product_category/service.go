package productcategory

import (
	"context"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
)

// Service interface contain functions
type Service interface {
	Create(context.Context, requestModel.CreateProductCategoryRequest) (*responseModel.CreateProductCategoryResponse, error)
	GetProductCategory(context.Context, requestModel.GetProductCategoryRequest) (*responseModel.GetProductCategoryResponse, error)
	GetProductCategories(context.Context, requestModel.GetProductCategoriesRequest) (*responseModel.GetProductCategoriesResponse, error)
	UpdateProductCategory(context.Context, requestModel.UpdateProductCategoryRequest) (*responseModel.UpdateProductCategoryResponse, error)
}
