package response

import (
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
)

// CreatePhotoResponse struct
type CreatePhotoResponse struct {
	Photo         *pgModel.Photo `json:"photo,omitempty"`
	TransactionID *string        `json:"transactionID,omitempty"`
}

// GetPhotosResponse struct
type GetPhotosResponse struct {
	Photos []*pgModel.Photo `json:"photos,omitempty"`
}

// RemovePhotoResponse struct
type RemovePhotoResponse struct {
	TransactionID *string `json:"transactionID,omitempty"`
}

// GetPhotoResponse struct
type GetPhotoResponse struct {
	Photo *pgModel.Photo `json:"photo,omitempty"`
}
