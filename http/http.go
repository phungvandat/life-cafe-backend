package http

import (
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	"github.com/phungvandat/life-cafe-backend/http/encode"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
	categoryRoute "github.com/phungvandat/life-cafe-backend/http/route/category"
	orderRoute "github.com/phungvandat/life-cafe-backend/http/route/order"
	productRoute "github.com/phungvandat/life-cafe-backend/http/route/product"
	publicResourceRoute "github.com/phungvandat/life-cafe-backend/http/route/public_resource"
	uploadRoute "github.com/phungvandat/life-cafe-backend/http/route/upload"
	userRoute "github.com/phungvandat/life-cafe-backend/http/route/user"
	"github.com/phungvandat/life-cafe-backend/util/helper"
)

func NewHTTPHandler(middlewares middlewares.Middlewares, endpoints endpoints.Endpoints,
	logger log.Logger,
	useCORS bool,
) http.Handler {
	r := chi.NewRouter()

	if useCORS {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})
		r.Use(cors.Handler)
	}

	options := []httpTransport.ServerOption{
		// Verify token jwt option
		httpTransport.ServerBefore(helper.VerifyToken),
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(encode.EncodeError),
	}

	r.Get("/", httpTransport.NewServer(
		endpoints.Index,
		httpTransport.NopRequestDecoder,
		httpTransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/users", userRoute.UserRoute(endpoints, middlewares, options))
	r.Route("/upload", uploadRoute.UploadRoute(endpoints, middlewares, options))
	r.Route("/categories", categoryRoute.CategoryRoute(endpoints, middlewares, options))
	r.Route("/products", productRoute.ProductRoute(endpoints, middlewares, options))
	r.Route("/orders", orderRoute.OrderRoute(endpoints, middlewares, options))
	r.Route("/public", publicResourceRoute.PublicResourceRoute(endpoints, middlewares, options))

	return r
}
