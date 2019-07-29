package user

import (
	"context"
	"strconv"
	"strings"

	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/constants"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
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

func (mw validationMiddleware) RollbackTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.RollbackTransaction(ctx, transactionID)
}

func (mw validationMiddleware) CommitTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.CommitTransaction(ctx, transactionID)
}

func (mw validationMiddleware) Create(ctx context.Context, req requestModel.CreateUserRequest) (*responseModel.CreateUserResponse, error) {
	if req.Username == "" {
		return nil, errors.MissingUsernameError
	}
	if !regex.IsUsernameValid(req.Username) {
		return nil, errors.InvalidUsernameError
	}

	if strings.Trim(req.Fullname, " ") == "" {
		return nil, errors.MissingFullnameError
	}

	if (req.Role == "master" || req.Role == "admin") && strings.Trim(req.Password, " ") == "" {
		return nil, errors.MissingPasswordError
	}

	if strings.Trim(req.Role, " ") == "" {
		return nil, errors.MissingRoleError
	}

	if _, check := constants.UserRole[req.Role]; !check {
		return nil, errors.InvalidRoleError
	}

	if req.PhoneNumber == "" {
		return nil, errors.UserPhoneNumberIsRequiredError
	}

	if !regex.IsPhoneNumberValid(req.PhoneNumber) {
		return nil, errors.InvalidUserPhoneNumberError
	}

	return mw.Service.Create(ctx, req)
}

func (mw validationMiddleware) LogIn(ctx context.Context, req requestModel.UserLogInRequest) (*responseModel.UserLogInResponse, error) {
	if req.Username == "" {
		return nil, errors.MissingUsernameError
	}

	if req.Password == "" {
		return nil, errors.MissingPasswordError
	}

	return mw.Service.LogIn(ctx, req)
}

func (mw validationMiddleware) GetUser(ctx context.Context, req requestModel.GetUserRequest) (*responseModel.GetUserResponse, error) {
	if _, err := pgModel.UUIDFromString(req.ParamUserID); err != nil {
		return nil, errors.InvalidUserIDError
	}

	return mw.Service.GetUser(ctx, req)
}

func (mw validationMiddleware) GetUsers(ctx context.Context, req requestModel.GetUsersRequest) (*responseModel.GetUsersResponse, error) {
	if _, err := strconv.ParseInt(req.Skip, 10, 32); req.Skip != "" && err != nil {
		return nil, errors.InvalidSkipError
	}

	if _, err := strconv.ParseInt(req.Limit, 10, 32); req.Limit != "" && err != nil {
		return nil, errors.InvalidLimitError
	}

	if _, check := constants.UserRole[req.Role]; req.Role != "" && !check {
		return nil, errors.InvalidRoleError
	}
	return mw.Service.GetUsers(ctx, req)
}
