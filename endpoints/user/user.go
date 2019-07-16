package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	requesetModel "github.com/phungvandat/life-cafe-backend/model/request"
	"github.com/phungvandat/life-cafe-backend/service"
)

// MakeCreateEndpoint make endpoint for create user
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.CreateUserRequest)
		res, err := s.UserService.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func MakeLogInEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.UserLogInRequest)
		res, err := s.UserService.LogIn(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
