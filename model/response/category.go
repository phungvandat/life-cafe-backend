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
	*pgModel.Category
	ChildrenCategories []*Category `json:"childrenCategories,omitempty"`
	ParentCategory     *Category   `json:"parentCategory,omitempty"`
}

// UpdateCategoryResponse struct
type UpdateCategoryResponse struct {
	Category *Category `json:"category,omitempty"`
}
