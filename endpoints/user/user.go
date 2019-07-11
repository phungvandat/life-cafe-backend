package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/phungvandat/life-cafe-backend/domain"
	"github.com/phungvandat/life-cafe-backend/service"
)

// MakeCreateEndpoint make endpoint for create user
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(CreateRequest)
			user = domain.User{
				Username: req.Username,
				Fullname: req.Fullname,
				Password: req.Password,
				Role:     req.Role,
			}
		)
		userResponse, err := s.UserService.Create(ctx, user)
		if err != nil {
			return nil, err
		}
		return CreateResponse{userResponse}, nil
	}
}

func MakeLogInEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(LogInRequest)
			user = domain.User{
				Username: req.Username,
				Password: req.Password,
			}
		)
		userResponse, token, err := s.UserService.LogIn(ctx, user)
		if err != nil {
			return nil, err
		}
		return LogInResponse{
			User:  userResponse,
			Token: token,
		}, nil
	}
}
