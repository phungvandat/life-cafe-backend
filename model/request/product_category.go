package request

// CreateCategoryRequest struct
type CreateCategoryRequest struct {
	Name             string `json:"name,omitempty"`
	Photo            string `json:"photo,omitempty"`
	ParentCategoryID string `json:"parent_category_id,omitempty"`
	Slug             string `json:"slug,omitempty"`
}

// GetCategoryRequest struct
type GetCategoryRequest struct {
	ParamCategoryID string `json:"category_id,omitempty"`
}

// GetCategoriesRequest struct
type GetCategoriesRequest struct {
	Skip  string `json:"skip,omitempty"`
	Limit string `json:"limit,omitempty"`
}

// UpdateCategoryRequest struct
type UpdateCategoryRequest struct {
	ParamCategoryID  string `json:"category_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Photo            string `json:"photo,omitempty"`
	ParentCategoryID string `json:"parent_category_id,omitempty"`
	Slug             string `json:"slug,omitempty"`
}
