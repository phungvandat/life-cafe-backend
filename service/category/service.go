package category

import (
	"context"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
)

// Service interface contain functions
type Service interface {
	Create(context.Context, requestModel.CreateCategoryRequest) (*responseModel.CreateCategoryResponse, error)
	GetCategory(context.Context, requestModel.GetCategoryRequest) (*responseModel.GetCategoryResponse, error)
	GetCategories(context.Context, requestModel.GetCategoriesRequest) (*responseModel.GetCategoriesResponse, error)
	UpdateCategory(context.Context, requestModel.UpdateCategoryRequest) (*responseModel.UpdateCategoryResponse, error)
}
