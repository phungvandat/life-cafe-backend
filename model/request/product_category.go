package request

// CreateProductCategoryRequest struct
type CreateProductCategoryRequest struct {
	Name             string `json:"name,omitempty"`
	Photo            string `json:"photo,omitempty"`
	ParentCategoryID string `json:"parent_category_id,omitempty"`
	Slug             string `json:"slug,omitempty"`
}

// GetProductCategoryRequest struct
type GetProductCategoryRequest struct {
	ParamProductCategoryID string `json:"product_category_id,omitempty"`
}

// GetProductCategoriesRequest struct
type GetProductCategoriesRequest struct {
	Skip  string `json:"skip,omitempty"`
	Limit string `json:"limit,omitempty"`
}

// UpdateProductCategoryRequest struct
type UpdateProductCategoryRequest struct {
	ParamProductCategoryID string `json:"product_category_id,omitempty"`
	Name                   string `json:"name,omitempty"`
	Photo                  string `json:"photo,omitempty"`
	ParentCategoryID       string `json:"parent_category_id,omitempty"`
	Slug                   string `json:"slug,omitempty"`
}
