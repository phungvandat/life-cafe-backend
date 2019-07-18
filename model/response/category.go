package response

import (
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
)

// CreateCategoryResponse struct
type CreateCategoryResponse struct {
	Category *pgModel.Category `json:"category,omitempty"`
}

// GetCategoryResponse struct
type GetCategoryResponse struct {
	Category *Category `json:"category,omitempty"`
}

// GetCategoriesResponse struct
type GetCategoriesResponse struct {
	Categories []*Category `json:"categories"`
}

// Category struct
type Category struct {
	ParentCategory *pgModel.Category `json:"parent_category,omitempty"`
	*pgModel.Category
}

// UpdateCategoryResponse struct
type UpdateCategoryResponse struct {
	Category *Category `json:"category,omitempty"`
}
