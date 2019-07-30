package user

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	userDecode "github.com/phungvandat/life-cafe-backend/http/decode/json/user"
	"github.com/phungvandat/life-cafe-backend/http/encode"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
)

// UserRoute func
func UserRoute(endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption) func(r chi.Router) {
	return func(r chi.Router) {
		// Create user
		r.Post("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin((endpoints.UserEndpoint.CreateUser)),
			userDecode.CreateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)

		// User log in
		r.Post("/log-in", httptransport.NewServer(
			endpoints.UserEndpoint.UserLogin,
			userDecode.LogInRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get users
		r.Get("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin((endpoints.UserEndpoint.GetUsers)),
			userDecode.GetUsersRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Get user
		r.Get("/{userID}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthUser(endpoints.UserEndpoint.GetUser),
			userDecode.GetUserRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Update user
		r.Put("/{userID}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthUser(
				endpoints.UserEndpoint.UpdateUser,
			),
			userDecode.UpdateUserRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
