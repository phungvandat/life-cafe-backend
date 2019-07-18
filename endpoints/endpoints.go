package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phungvandat/life-cafe-backend/endpoints/index"
	"github.com/phungvandat/life-cafe-backend/endpoints/product"
	productcategory "github.com/phungvandat/life-cafe-backend/endpoints/product_category"
	"github.com/phungvandat/life-cafe-backend/endpoints/upload"
	"github.com/phungvandat/life-cafe-backend/endpoints/user"
	"github.com/phungvandat/life-cafe-backend/service"
)

// Endpoints struct
type Endpoints struct {
	Index                   endpoint.Endpoint
	UserEndpoint            user.UserEndpoint
	UploadEndpoint          upload.UploadEndpoint
	ProductCategoryEndpoint productcategory.ProductCategoryEndpoint
	ProductEndpoint         product.ProductEndpoint
}

// MakeServerEndpoints func
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Index: index.MakeIndexEndpoints(),
		UserEndpoint: user.UserEndpoint{
			CreateUser: user.MakeCreateEndpoint(s),
			UserLogin:  user.MakeLogInEndpoint(s),
		},
		UploadEndpoint: upload.UploadEndpoint{
			UploadImages: upload.MakeUploadImages(s),
		},
		ProductCategoryEndpoint: productcategory.ProductCategoryEndpoint{
			CreateProductCategory: productcategory.MakeCreateEndpoint(s),
			GetProductCategory:    productcategory.MakeGetProductCategoryEndpoint(s),
			GetProductCategories:  productcategory.MakeGetProductCategoriesEndpoint(s),
			UpdateProductCategory: productcategory.MakeUpdateProductCategoryEndpoint(s),
		},
		ProductEndpoint: product.ProductEndpoint{
			CreateProductEndpoint: product.MakeCreateProductEndpoint(s),
			GetProductEndpoint:    product.MakeGetProductEndpoint(s),
			GetProductsEndpoint:   product.MakeGetProductsEndpoint(s),
			UpdateProductEndpoint: product.MakeUpdateProductEndpoint(s),
		},
	}
}
