package productcategory

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	decode "github.com/phungvandat/life-cafe-backend/http/decode/json/product_category"
	"github.com/phungvandat/life-cafe-backend/http/encode"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
)

// ProductCategoryRoute func
func ProductCategoryRoute(endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption,
) func(r chi.Router) {
	return func(r chi.Router) {
		// Create product category
		r.Post("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(endpoints.ProductCategoryEndpoint.CreateProductCategory),
			decode.CreateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get product category
		r.Get("/{product_category_id}", httptransport.NewServer(
			endpoints.ProductCategoryEndpoint.GetProductCategory,
			decode.GetProductCategoryRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get product categories
		r.Get("/", httptransport.NewServer(
			endpoints.ProductCategoryEndpoint.GetProductCategories,
			decode.GetProductCategoriesRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Update product category
		r.Put("/{product_category_id}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(endpoints.ProductCategoryEndpoint.UpdateProductCategory),
			decode.UpdateProductCategoryRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
