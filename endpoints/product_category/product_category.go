package productcategory

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	requesetModel "github.com/phungvandat/life-cafe-backend/model/request"
	"github.com/phungvandat/life-cafe-backend/service"
)

// MakeCreateEndpoint make endpoint for create user
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.CreateProductCategoryRequest)
		res, err := s.ProductCategoryService.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeGetProductCategoryEndpoint func
func MakeGetProductCategoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetProductCategoryRequest)
		res, err := s.ProductCategoryService.GetProductCategory(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeGetProductCategoriesEndpoint func
func MakeGetProductCategoriesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetProductCategoriesRequest)

		res, err := s.ProductCategoryService.GetProductCategories(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeUpdateProductCategoryEndpoint func
func MakeUpdateProductCategoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.UpdateProductCategoryRequest)

		res, err := s.ProductCategoryService.UpdateProductCategory(ctx, req)

		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
