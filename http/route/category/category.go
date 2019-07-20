package category

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	decode "github.com/phungvandat/life-cafe-backend/http/decode/json/category"
	"github.com/phungvandat/life-cafe-backend/http/encode"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
)

// CategoryRoute func
func CategoryRoute(endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption,
) func(r chi.Router) {
	return func(r chi.Router) {
		// Create category
		r.Post("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(endpoints.CategoryEndpoint.CreateCategory),
			decode.CreateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get category
		r.Get("/{category_id}", httptransport.NewServer(
			endpoints.CategoryEndpoint.GetCategory,
			decode.GetCategoryRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get product categories
		r.Get("/", httptransport.NewServer(
			endpoints.CategoryEndpoint.GetCategories,
			decode.GetCategoriesRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Update category
		r.Put("/{category_id}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(endpoints.CategoryEndpoint.UpdateCategory),
			decode.UpdateCategoryRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
