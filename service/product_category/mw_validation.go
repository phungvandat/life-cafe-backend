package productcategory

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	domainModel "github.com/phungvandat/life-cafe-backend/model/domain"
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

func (mw validationMiddleware) Create(ctx context.Context, req requestModel.CreateProductCategoryRequest) (*responseModel.CreateProductCategoryResponse, error) {
	if strings.Trim(req.Name, " ") == "" {
		return nil, ProductCategoryNameIsRequiredError
	}

	if _, err := domainModel.UUIDFromString(req.ParentCategoryID); req.ParentCategoryID != "" && err != nil {
		return nil, InvalidParentProductCategoryIDTypeError
	}

	if strings.Trim(req.Slug, " ") == "" {
		return nil, ProductCategorySlugIsRequiredError
	}

	slugRegex, _ := regexp.Compile(regex.SlugRegex)

	if !slugRegex.MatchString(req.Slug) {
		return nil, InvalidProductCategorySlugError
	}

	return mw.Service.Create(ctx, req)
}

func (mw validationMiddleware) GetProductCategory(ctx context.Context, req requestModel.GetProductCategoryRequest) (*responseModel.GetProductCategoryResponse, error) {
	if _, err := domainModel.UUIDFromString(req.ParamProductCategoryID); req.ParamProductCategoryID != "" && err != nil {
		return nil, InvalidProductCategoryIDTypeError
	}
	return mw.Service.GetProductCategory(ctx, req)
}

func (mw validationMiddleware) GetProductCategories(ctx context.Context, req requestModel.GetProductCategoriesRequest) (*responseModel.GetProductCategoriesResponse, error) {
	if _, err := strconv.ParseInt(req.Skip, 10, 32); req.Skip != "" && err != nil {
		return nil, InvalidSkipError
	}

	if _, err := strconv.ParseInt(req.Limit, 10, 32); req.Limit != "" && err != nil {
		return nil, InvalidLimitError
	}
	return mw.Service.GetProductCategories(ctx, req)
}

func (mw validationMiddleware) UpdateProductCategory(ctx context.Context, req requestModel.UpdateProductCategoryRequest) (*responseModel.UpdateProductCategoryResponse, error) {
	if _, err := domainModel.UUIDFromString(req.ParamProductCategoryID); req.ParamProductCategoryID != "" && err != nil {
		return nil, InvalidProductCategoryIDTypeError
	}

	if req.Name != "" && strings.Trim(req.Name, " ") == "" {
		return nil, InvalidProductCategoryNameError
	}

	if _, err := domainModel.UUIDFromString(req.ParentCategoryID); req.ParentCategoryID != "" && err != nil {
		return nil, InvalidParentProductCategoryIDTypeError
	}

	slugRegex, _ := regexp.Compile(regex.SlugRegex)

	if req.Slug != "" && !slugRegex.MatchString(req.Slug) {
		return nil, InvalidProductCategorySlugError
	}

	return mw.Service.UpdateProductCategory(ctx, req)
}
