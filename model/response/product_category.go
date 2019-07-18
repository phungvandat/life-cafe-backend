package response

import (
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
)

// CreateProductCategoryResponse struct
type CreateProductCategoryResponse struct {
	ProductCategory *pgModel.ProductCategory `json:"product_category,omitempty"`
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
	ParentCategory *pgModel.ProductCategory `json:"parent_category,omitempty"`
	*pgModel.ProductCategory
}

// UpdateProductCategoryResponse struct
type UpdateProductCategoryResponse struct {
	ProductCategory *ProductCategory `json:"product_category,omitempty"`
}
