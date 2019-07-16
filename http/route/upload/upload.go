package upload

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	uploadDecode "github.com/phungvandat/life-cafe-backend/http/decode/json/upload"
	"github.com/phungvandat/life-cafe-backend/http/encode"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
)

func UploadRoute(endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption) func(r chi.Router) {
	return func(r chi.Router) {
		// Upload image
		r.Post("/images", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthUser((endpoints.UploadEndpoint.UploadImages)),
			uploadDecode.UploadImagesRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
