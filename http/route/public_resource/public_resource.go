package public

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	uploadDecode "github.com/phungvandat/life-cafe-backend/http/decode/json/upload"
	"github.com/phungvandat/life-cafe-backend/http/encode"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
)

// PublicResourceRoute func
func PublicResourceRoute(endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption) func(r chi.Router) {
	return func(r chi.Router) {
		// Upload image
		r.Get("/images/{file_path}", httptransport.NewServer(
			endpoints.UploadEndpoint.GetImageFile,
			uploadDecode.GetImageFileRequest,
			encode.EncodeFileResponse,
			options...,
		).ServeHTTP)
	}
}
