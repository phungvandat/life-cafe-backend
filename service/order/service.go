package order

import (
	"context"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
)

// Service interface contain functions
type Service interface {
	RollbackTransaction(context.Context, string) error
	CommitTransaction(context.Context, string) error
	CreateOrder(context.Context, requestModel.CreateOrderRequest) (*responseModel.CreateOrderResponse, error)
	GetOrder(context.Context, requestModel.GetOrderRequest) (*responseModel.GetOrderResponse, error)
	GetOrders(context.Context, requestModel.GetOrdersRequest) (*responseModel.GetOrdersResponse, error)
	UpdateOrder(context.Context, requestModel.UpdateOrderRequest) (*responseModel.UpdateOrderResponse, error)
}
