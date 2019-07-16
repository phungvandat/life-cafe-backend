package user

import (
	"context"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
)

// Service interface contain functions
type Service interface {
	Create(context.Context, requestModel.CreateUserRequest) (*responseModel.CreateUserResponse, error)
	LogIn(context.Context, requestModel.UserLogInRequest) (*responseModel.UserLogInResponse, error)
	CreateMaster(context.Context) error
}
