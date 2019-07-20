package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	requesetModel "github.com/phungvandat/life-cafe-backend/model/request"
	"github.com/phungvandat/life-cafe-backend/service"
)

// MakeCreateOrderEnpoint func
func MakeCreateOrderEnpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.CreateOrderRequest)

		res, err := s.OrderService.CreateOrder(ctx, req)

		var errTransaction error
		var transactionID *string
		if res != nil && res.TransactionID != nil {
			transactionID = res.TransactionID
		}
		if err != nil {
			if transactionID != nil {
				errTransaction = s.OrderService.RollbackTransaction(ctx, *transactionID)
			}
			return nil, err
		}
		if transactionID != nil {
			errTransaction = s.OrderService.CommitTransaction(ctx, *transactionID)
			res.TransactionID = nil
		}
		if errTransaction != nil {
			var log log.Logger
			log.Log("Transaction create order failure by error ", errTransaction)
		}
		return res, nil
	}
}

// MakeGetOrderEndpoint func
func MakeGetOrderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetOrderRequest)

		res, err := s.OrderService.GetOrder(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeGetOrdersEndpoint func
func MakeGetOrdersEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetOrdersRequest)
		res, err := s.OrderService.GetOrders(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeUpdateOrderEndpoint func
func MakeUpdateOrderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.UpdateOrderRequest)

		res, err := s.OrderService.UpdateOrder(ctx, req)

		var errTransaction error
		var transactionID *string
		if res != nil && res.TransactionID != nil {
			transactionID = res.TransactionID
		}

		if err != nil {
			if transactionID != nil {
				errTransaction = s.OrderService.RollbackTransaction(ctx, *transactionID)
			}
			return nil, err
		}

		if transactionID != nil {
			errTransaction = s.OrderService.CommitTransaction(ctx, *transactionID)
			res.TransactionID = nil
		}

		if errTransaction != nil {
			var log log.Logger
			log.Log("Transaction update order failure by error ", errTransaction)
		}

		return res, nil
	}
}
