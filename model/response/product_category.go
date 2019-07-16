package response

import (
	domainModel "github.com/phungvandat/life-cafe-backend/model/domain"
)

// CreateProductCategoryResponse struct
type CreateProductCategoryResponse struct {
	ProductCategory *domainModel.ProductCategory `json:"product_category,omitempty"`
}

// GetProductCategoryResponse struct
type GetProductCategoryResponse struct {
	ProductCategory *ProductCategory `json:"product_category,omitempty"`
}

// GetProductCategoriesResponse struct
type GetProductCategoriesResponse struct {
	ProductCategories []*ProductCategory `json:"product_categories"`
}

// ProductCategory struct
type ProductCategory struct {
	ParentCategory *domainModel.ProductCategory `json:"parent_category,omitempty"`
	*domainModel.ProductCategory
}

// UpdateProductCategoryResponse struct
type UpdateProductCategoryResponse struct {
	ProductCategory *ProductCategory `json:"product_category,omitempty"`
}
