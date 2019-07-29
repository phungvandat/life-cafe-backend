package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

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

// GetUsersRequest func
func GetUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetUsersRequest

	req.Skip = r.URL.Query().Get("skip")
	req.Limit = r.URL.Query().Get("limit")
	req.Fullname = r.URL.Query().Get("fullname")
	req.PhoneNumber = r.URL.Query().Get("phoneNumber")
	req.AlwaysPhone = r.URL.Query().Get("alwaysPhone")
	req.Role = r.URL.Query().Get("role")

	return req, nil
}

// GetUserRequest func
func GetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetUserRequest
	req.ParamUserID = chi.URLParam(r, "userID")

	return req, nil
}
