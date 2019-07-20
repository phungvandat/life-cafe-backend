package service

import (
	"github.com/phungvandat/life-cafe-backend/service/auth"
	"github.com/phungvandat/life-cafe-backend/service/category"
	"github.com/phungvandat/life-cafe-backend/service/order"
	"github.com/phungvandat/life-cafe-backend/service/product"
	"github.com/phungvandat/life-cafe-backend/service/upload"
	"github.com/phungvandat/life-cafe-backend/service/user"
)

// Service define list of all services in project
type Service struct {
	UserService     user.Service
	AuthService     auth.Service
	UploadService   upload.Service
	CategoryService category.Service
	ProductService  product.Service
	OrderService    order.Service
}
