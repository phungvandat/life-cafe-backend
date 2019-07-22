package order

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
)

// CreateOrderRequest func
func CreateOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// GetOrderRequest func
func GetOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetOrderRequest
	orderID := chi.URLParam(r, "orderID")
	req.ParamOrderID = orderID
	return req, nil
}

// GetOrdersRequest func
func GetOrdersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetOrdersRequest

	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")

	req.Skip = skip
	req.Limit = limit

	return req, nil
}

// UpdateOrderRequest func
func UpdateOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.UpdateOrderRequest
	orderID := chi.URLParam(r, "orderID")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, err
	}
	req.ParamOrderID = orderID
	return req, nil
}
