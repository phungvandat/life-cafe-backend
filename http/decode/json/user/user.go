package user

import (
	"context"
	"encoding/json"
	"net/http"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// LogInRequest func
func LogInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.UserLogInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
