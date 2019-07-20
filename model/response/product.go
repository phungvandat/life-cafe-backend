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
	SubPhotos  []string            `json:"sub_photos"`
	Categories []*pgModel.Category `json:"categories"`
}

// GetProductResponse struct
type GetProductResponse struct {
	Product *Product `json:"product,omitempty"`
}

// GetProductsResponse struct
type GetProductsResponse struct {
	Products []*Product `json:"products"`
}

// UpdateProductResponse struct
type UpdateProductResponse struct {
	Product       *Product `json:"product,omitempty"`
	TransactionID *string  `json:"transaction_id,omitempty"`
}
