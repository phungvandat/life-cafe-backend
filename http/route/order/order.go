package order

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	decode "github.com/phungvandat/life-cafe-backend/http/decode/json/order"
	"github.com/phungvandat/life-cafe-backend/http/encode"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
)

// OrderRoute func
func OrderRoute(endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption,
) func(r chi.Router) {
	return func(r chi.Router) {
		// Create order
		r.Post("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(
				endpoints.OrderEndpoint.CreateOrderEndpoint,
			),
			decode.CreateOrderRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get order
		r.Get("/{order_id}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthUser(
				endpoints.OrderEndpoint.GetOrderEndpoint,
			),
			decode.GetOrderRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get orders
		r.Get("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(
				endpoints.OrderEndpoint.GetOrdersEndpoint,
			),
			decode.GetOrdersRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Update order
		r.Put("/{order_id}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(
				endpoints.OrderEndpoint.UpdateOrderEndpoint,
			),
			decode.UpdateOrderRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
