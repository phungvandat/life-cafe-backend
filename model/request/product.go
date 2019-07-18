package request

// CreateProductRequest struct
type CreateProductRequest struct {
	Name        string   `json:"name,omitempty"`
	MainPhoto   string   `json:"main_photo,omitempty"`
	SubPhotos   []string `json:"sub_photos,omitempty"`
	Quantity    int      `json:"quantity,omitempty"`
	Price       int      `json:"price,omitempty"`
	Flag        int      `json:"flag,omitempty"`
	Slug        string   `json:"slug,omitempty"`
	Barcode     string   `json:"barcode,omitempty"`
	CategoryIDs []string `json:"category_ids,omitempty"`
	Description string   `json:"description,omitempty"`
}

// GetProductRequest struct
type GetProductRequest struct {
	ParamProductID string `json:"product_id,omitempty"`
}

// GetProductsRequest struct
type GetProductsRequest struct {
	Skip  string `json:"skip,omitempty"`
	Limit string `json:"limit,omitempty"`
}

// UpdateProductRequest struct
type UpdateProductRequest struct {
	ParamProductID string   `json:"product_id,omitempty"`
	Name           string   `json:"name,omitempty"`
	MainPhoto      string   `json:"main_photo,omitempty"`
	SubPhotos      []string `json:"sub_photos,omitempty"`
	Price          int      `json:"price,omitempty"`
	Flag           int      `json:"flag,omitempty"`
	Slug           string   `json:"slug,omitempty"`
	Barcode        string   `json:"barcode,omitempty"`
	CategoryIDs    []string `json:"category_ids,omitempty"`
	Description    string   `json:"description,omitempty"`
}
