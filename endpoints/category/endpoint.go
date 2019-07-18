package category

import (
	"github.com/go-kit/kit/endpoint"
)

type CategoryEndpoint struct {
	CreateCategory endpoint.Endpoint
	GetCategory    endpoint.Endpoint
	GetCategories  endpoint.Endpoint
	UpdateCategory endpoint.Endpoint
}
