package product

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	decode "github.com/phungvandat/life-cafe-backend/http/decode/json/product"
	"github.com/phungvandat/life-cafe-backend/http/encode"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
)

// ProductRoute func
func ProductRoute(endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption,
) func(r chi.Router) {
	return func(r chi.Router) {
		// Create product
		r.Post("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(endpoints.ProductEndpoint.CreateProductEndpoint),
			decode.CreateProductRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get product
		r.Get("/{product_id}", httptransport.NewServer(
			endpoints.ProductEndpoint.GetProductEndpoint,
			decode.GetProductRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get products
		r.Get("/", httptransport.NewServer(
			endpoints.ProductEndpoint.GetProductsEndpoint,
			decode.GetProductsRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Update product
		r.Put("/{product_id}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(endpoints.ProductEndpoint.UpdateProductEndpoint),
			decode.UpdateProductRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
