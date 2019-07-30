package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	requesetModel "github.com/phungvandat/life-cafe-backend/model/request"
	"github.com/phungvandat/life-cafe-backend/service"
)

// MakeCreateEndpoint make endpoint for create user
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.CreateUserRequest)
		res, err := s.UserService.Create(ctx, req)
		var errTransaction error
		var transactionID *string
		if res != nil && res.TransactionID != nil {
			transactionID = res.TransactionID
		}
		if err != nil {
			if transactionID != nil {
				errTransaction = s.UserService.RollbackTransaction(ctx, *transactionID)
			}
			return nil, err
		}
		if transactionID != nil {
			errTransaction = s.UserService.CommitTransaction(ctx, *transactionID)
			res.TransactionID = nil
		}
		if errTransaction != nil {
			var log log.Logger
			log.Log("Transaction create user failure by error ", errTransaction)
		}
		return res, nil
	}
}

// MakeLogInEndpoint func
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

// MakeGetUsersEndpoint func
func MakeGetUsersEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetUsersRequest)
		res, err := s.UserService.GetUsers(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeGetUserEndpoint func
func MakeGetUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetUserRequest)

		res, err := s.UserService.GetUser(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeUpdateUserEndpoint func
func MakeUpdateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.UpdateUserRequest)

		res, err := s.UserService.UpdateUser(ctx, req)

		var errTransaction error
		var transactionID *string

		if res != nil && res.TransactionID != nil {
			transactionID = res.TransactionID
		}
		if err != nil {
			if transactionID != nil {
				errTransaction = s.UserService.RollbackTransaction(ctx, *transactionID)
			}
			return nil, err
		}
		if transactionID != nil {
			errTransaction = s.UserService.CommitTransaction(ctx, *transactionID)
			res.TransactionID = nil
		}
		if errTransaction != nil {
			var log log.Logger
			log.Log("Transaction update user failure by error ", errTransaction)
		}

		return res, nil
	}
}
