package service

import (
	"github.com/phungvandat/life-cafe-backend/service/auth"
	"github.com/phungvandat/life-cafe-backend/service/product"
	productcategory "github.com/phungvandat/life-cafe-backend/service/product_category"
	"github.com/phungvandat/life-cafe-backend/service/upload"
	"github.com/phungvandat/life-cafe-backend/service/user"
)

// Service define list of all services in project
type Service struct {
	UserService            user.Service
	AuthService            auth.Service
	UploadService          upload.Service
	ProductCategoryService productcategory.Service
	ProductService         product.Service
}
