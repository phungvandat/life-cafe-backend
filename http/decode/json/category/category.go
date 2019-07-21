package category

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.CreateCategoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

//GetCategoryRequest func
func GetCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetCategoryRequest

	categoryID := chi.URLParam(r, "category_id")
	req.ParamCategoryID = categoryID

	return req, nil
}

// GetCategoriesRequest func
func GetCategoriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.GetCategoriesRequest

	req.Skip = r.URL.Query().Get("skip")
	req.Limit = r.URL.Query().Get("limit")
	req.Slug = r.URL.Query().Get("slug")
	req.ParentCategoryExist = r.URL.Query().Get("parent_category_exist")

	return req, nil
}

// UpdateCategoryRequest func
func UpdateCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.UpdateCategoryRequest
	categoryID := chi.URLParam(r, "category_id")
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}
	req.ParamCategoryID = categoryID
	return req, nil
}
