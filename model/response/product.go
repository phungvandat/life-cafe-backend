package response

import (
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
)

// CreateProductResponse struct
type CreateProductResponse struct {
	Product       *Product `json:"product,omitempty"`
	TransactionID *string  `json:"transactionID,omitempty"`
}

// Product struct
type Product struct {
	*pgModel.Product
	SubPhotos  []*pgModel.Photo    `json:"subPhotos"`
	Categories []*pgModel.Category `json:"categories"`
}

// GetProductResponse struct
type GetProductResponse struct {
	Product *Product `json:"product,omitempty"`
}

// GetProductsResponse struct
type GetProductsResponse struct {
	Products []*Product `json:"products"`
	Total    int        `json:"total,omitempty"`
}

// UpdateProductResponse struct
type UpdateProductResponse struct {
	Product       *Product `json:"product,omitempty"`
	TransactionID *string  `json:"transactionID,omitempty"`
}
