package category

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	requesetModel "github.com/phungvandat/life-cafe-backend/model/request"
	"github.com/phungvandat/life-cafe-backend/service"
)

// MakeCreateEndpoint make endpoint for create user
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.CreateCategoryRequest)
		res, err := s.CategoryService.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeGetCategoryEndpoint func
func MakeGetCategoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetCategoryRequest)
		res, err := s.CategoryService.GetCategory(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeGetCategoriesEndpoint func
func MakeGetCategoriesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetCategoriesRequest)

		res, err := s.CategoryService.GetCategories(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeUpdateCategoryEndpoint func
func MakeUpdateCategoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.UpdateCategoryRequest)

		res, err := s.CategoryService.UpdateCategory(ctx, req)

		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
