package category

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/regex"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware for validation purposes
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, req requestModel.CreateCategoryRequest) (*responseModel.CreateCategoryResponse, error) {
	if strings.Trim(req.Name, " ") == "" {
		return nil, CategoryNameIsRequiredError
	}

	if _, err := pgModel.UUIDFromString(req.ParentCategoryID); req.ParentCategoryID != "" && err != nil {
		return nil, InvalidParentCategoryIDTypeError
	}

	if strings.Trim(req.Slug, " ") == "" {
		return nil, CategorySlugIsRequiredError
	}

	slugRegex, _ := regexp.Compile(regex.SlugRegex)

	if !slugRegex.MatchString(req.Slug) {
		return nil, InvalidCategorySlugError
	}

	return mw.Service.Create(ctx, req)
}

func (mw validationMiddleware) GetCategory(ctx context.Context, req requestModel.GetCategoryRequest) (*responseModel.GetCategoryResponse, error) {
	if _, err := pgModel.UUIDFromString(req.ParamCategoryID); req.ParamCategoryID != "" && err != nil {
		return nil, InvalidCategoryIDTypeError
	}
	return mw.Service.GetCategory(ctx, req)
}

func (mw validationMiddleware) GetCategories(ctx context.Context, req requestModel.GetCategoriesRequest) (*responseModel.GetCategoriesResponse, error) {
	if _, err := strconv.ParseInt(req.Skip, 10, 32); req.Skip != "" && err != nil {
		return nil, InvalidSkipError
	}

	if _, err := strconv.ParseInt(req.Limit, 10, 32); req.Limit != "" && err != nil {
		return nil, InvalidLimitError
	}
	return mw.Service.GetCategories(ctx, req)
}

func (mw validationMiddleware) UpdateCategory(ctx context.Context, req requestModel.UpdateCategoryRequest) (*responseModel.UpdateCategoryResponse, error) {
	if _, err := pgModel.UUIDFromString(req.ParamCategoryID); req.ParamCategoryID != "" && err != nil {
		return nil, InvalidCategoryIDTypeError
	}

	if req.Name != "" && strings.Trim(req.Name, " ") == "" {
		return nil, InvalidCategoryNameError
	}

	if _, err := pgModel.UUIDFromString(req.ParentCategoryID); req.ParentCategoryID != "" && err != nil {
		return nil, InvalidParentCategoryIDTypeError
	}

	slugRegex, _ := regexp.Compile(regex.SlugRegex)

	if req.Slug != "" && !slugRegex.MatchString(req.Slug) {
		return nil, InvalidCategorySlugError
	}

	return mw.Service.UpdateCategory(ctx, req)
}
