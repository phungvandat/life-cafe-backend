package request

// CreatePhotoRequest struct
type CreatePhotoRequest struct {
	URL       string `json:"url,omitempty"`
	ProductID string `json:"productID,omitempty"`
	PhotoID   string `json:"photoID,omitempty"`
}

// GetPhotosRequest struct
type GetPhotosRequest struct {
	ProductID string `json:"productID,omitempty"`
	Skip      string `json:"skip,omitempty"`
	Limit     string `json:"limit,omitempty"`
}

// RemovePhotoRequest struct
type RemovePhotoRequest struct {
	ParamPhotoID string `json:"photoID,omitempty"`
	ProductID    string `json:"productID,omitempty"`
}

// GetPhotoRequest struct
type GetPhotoRequest struct {
	ParamPhotoID string `json:"photoID,omitempty"`
}
