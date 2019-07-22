package request

// CreateCategoryRequest struct
type CreateCategoryRequest struct {
	Name             string `json:"name,omitempty"`
	Photo            string `json:"photo,omitempty"`
	ParentCategoryID string `json:"parentCategoryID,omitempty"`
	Slug             string `json:"slug,omitempty"`
}

// GetCategoryRequest struct
type GetCategoryRequest struct {
	ParamCategoryID string `json:"categoryID,omitempty"`
}

// GetCategoriesRequest struct
type GetCategoriesRequest struct {
	Skip                string `json:"skip,omitempty"`
	Limit               string `json:"limit,omitempty"`
	Slug                string `json:"string,omitempty"`
	ParentCategoryExist string `json:"parentCategoryExist,omitempty"`
}

// UpdateCategoryRequest struct
type UpdateCategoryRequest struct {
	ParamCategoryID  string `json:"categoryID,omitempty"`
	Name             string `json:"name,omitempty"`
	Photo            string `json:"photo,omitempty"`
	ParentCategoryID string `json:"parentCategoryID,omitempty"`
	Slug             string `json:"slug,omitempty"`
}
