package response

import (
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
)

// CreateProductResponse struct
type CreateProductResponse struct {
	Product       *Product `json:"product,omitempty"`
	TransactionID *string  `json:"transaction_id,omitempty"`
}

// Product struct
type Product struct {
	*pgModel.Product
	SubPhotos  []string                   `json:"sub_photos,omitempty"`
	Categories []*pgModel.ProductCategory `json:"categories,omitempty"`
}

// GetProductResponse struct
type GetProductResponse struct {
	Product *Product `json:"product,omitempty"`
}

// GetProductsResponse struct
type GetProductsResponse struct {
	Products []*Product `json:"products,omitempty"`
}

// UpdateProductResponse struct
type UpdateProductResponse struct {
	Product       *Product `json:"product,omitempty"`
	TransactionID *string  `json:"transaction_id,omitempty"`
}
