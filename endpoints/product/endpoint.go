package product

import (
	"github.com/go-kit/kit/endpoint"
)

// ProductEndpoint struct
type ProductEndpoint struct {
	CreateProductEndpoint endpoint.Endpoint
	GetProductEndpoint    endpoint.Endpoint
	GetProductsEndpoint   endpoint.Endpoint
	UpdateProductEndpoint endpoint.Endpoint
}
