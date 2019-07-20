package order

import (
	"github.com/go-kit/kit/endpoint"
)

// OrderEndpoint struct
type OrderEndpoint struct {
	CreateOrderEndpoint endpoint.Endpoint
	GetOrderEndpoint    endpoint.Endpoint
	GetOrdersEndpoint   endpoint.Endpoint
	UpdateOrderEndpoint endpoint.Endpoint
}
