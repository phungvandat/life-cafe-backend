package user

import (
	"context"
	"regexp"
	"strings"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/constants"
	"github.com/phungvandat/life-cafe-backend/util/regex"
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

func (mw validationMiddleware) Create(ctx context.Context, req requestModel.CreateUserRequest) (*responseModel.CreateUserResponse, error) {
	if req.Username == "" {
		return nil, MissingUsernameError
	}
	usernameRegex, _ := regexp.Compile(regex.UsernameRegex)
	if !usernameRegex.MatchString(req.Username) {
		return nil, InvalidUsernameError
	}

	if strings.Trim(req.Fullname, " ") == "" {
		return nil, MissingFullnameError
	}

	if strings.Trim(req.Password, " ") == "" {
		return nil, MissingPasswordError
	}

	if strings.Trim(req.Role, " ") == "" {
		return nil, MissingRoleError
	}

	if _, check := constants.USER_ROLE[req.Role]; !check {
		return nil, InvalidRoleError
	}

	return mw.Service.Create(ctx, req)
}

func (mw validationMiddleware) LogIn(ctx context.Context, req requestModel.UserLogInRequest) (*responseModel.UserLogInResponse, error) {
	if req.Username == "" {
		return nil, MissingUsernameError
	}

	if req.Password == "" {
		return nil, MissingPasswordError
	}

	return mw.Service.LogIn(ctx, req)
}

func (mw validationMiddleware) RollbackTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.RollbackTransaction(ctx, transactionID)
}

func (mw validationMiddleware) CommitTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.CommitTransaction(ctx, transactionID)
}
