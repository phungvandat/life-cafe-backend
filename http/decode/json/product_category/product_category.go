package productcategory

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.CreateProductCategoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

//GetProductCategoryRequest func
func GetProductCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetProductCategoryRequest

	productCategoryID := chi.URLParam(r, "product_category_id")
	req.ParamProductCategoryID = productCategoryID

	return req, nil
}

// GetProductCategoriesRequest func
func GetProductCategoriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetProductCategoriesRequest
	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")

	req.Skip = skip
	req.Limit = limit
	return req, nil
}

// UpdateProductCategoryRequest func
func UpdateProductCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.UpdateProductCategoryRequest
	productCategoryID := chi.URLParam(r, "product_category_id")
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}
	req.ParamProductCategoryID = productCategoryID
	return req, nil
}
