package user

import (
	"context"
	"encoding/json"
	"net/http"

	userEndpoint "github.com/phungvandat/life-cafe-backend/endpoints/user"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// LogInRequest func
func LogInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userEndpoint.LogInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
