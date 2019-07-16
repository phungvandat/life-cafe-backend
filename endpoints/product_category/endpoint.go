package productcategory

import (
	"github.com/go-kit/kit/endpoint"
)

type ProductCategoryEndpoint struct {
	CreateProductCategory endpoint.Endpoint
	GetProductCategory    endpoint.Endpoint
	GetProductCategories  endpoint.Endpoint
	UpdateProductCategory endpoint.Endpoint
}
