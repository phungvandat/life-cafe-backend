package product

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
)

// CreateProductRequest func
func CreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// GetProductRequest func
func GetProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetProductRequest
	productID := chi.URLParam(r, "product_id")

	req.ParamProductID = productID
	return req, nil
}

// GetProductsRequest func
func GetProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetProductsRequest

	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")

	req.Skip = skip
	req.Limit = limit
	return req, nil
}

func UpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.UpdateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	productID := chi.URLParam(r, "product_id")

	req.ParamProductID = productID
	return req, nil
}
