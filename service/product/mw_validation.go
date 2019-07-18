package product

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/helper"
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

func (mw validationMiddleware) RollbackTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.RollbackTransaction(ctx, transactionID)
}

func (mw validationMiddleware) CommitTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.CommitTransaction(ctx, transactionID)
}

func (mw validationMiddleware) CreateProduct(ctx context.Context, req requestModel.CreateProductRequest) (*responseModel.CreateProductResponse, error) {
	if strings.Trim(req.Name, " ") == "" {
		return nil, ProductNameIsRequiredError
	}

	if req.MainPhoto == "" {
		return nil, ProductMainPhotoIsRequiredError
	}

	if req.Quantity < 0 {
		return nil, InvalidProductQuantityError
	}

	if req.CategoryIDs == nil || len(req.CategoryIDs) <= 0 {
		return nil, CategoryOfProductIsRequiredError
	}

	uniqueCategoryIDs := helper.UniqueStringArray(req.CategoryIDs)

	if len(req.CategoryIDs) != len(uniqueCategoryIDs) {
		return nil, DuplicateCategoryError
	}

	for _, categoryID := range req.CategoryIDs {
		if _, err := pgModel.UUIDFromString(categoryID); err != nil {
			return nil, InvalidCategoryIDError
		}
	}

	if req.Price < 0 {
		return nil, InvalidProductPriceError
	}

	if strings.Trim(req.Slug, " ") == "" {
		return nil, ProductSlugIsRequiredError
	}

	slugRegex, _ := regexp.Compile(regex.SlugRegex)

	if !slugRegex.MatchString(req.Slug) {
		return nil, InvalidProductSlugError
	}

	if len(req.SubPhotos) > 3 {
		return nil, InvalidSecondaryPhotoQuantityError
	}

	return mw.Service.CreateProduct(ctx, req)
}

func (mw validationMiddleware) GetProduct(ctx context.Context, req requestModel.GetProductRequest) (*responseModel.GetProductResponse, error) {
	if _, err := pgModel.UUIDFromString(req.ParamProductID); req.ParamProductID != "" && err != nil {
		return nil, InvalidProductIDTypeError
	}
	return mw.Service.GetProduct(ctx, req)
}

func (mw validationMiddleware) GetProducts(ctx context.Context, req requestModel.GetProductsRequest) (*responseModel.GetProductsResponse, error) {
	if _, err := strconv.ParseInt(req.Skip, 10, 32); req.Skip != "" && err != nil {
		return nil, InvalidSkipError
	}

	if _, err := strconv.ParseInt(req.Limit, 10, 32); req.Limit != "" && err != nil {
		return nil, InvalidLimitError
	}

	return mw.Service.GetProducts(ctx, req)
}

func (mw validationMiddleware) UpdateProduct(ctx context.Context, req requestModel.UpdateProductRequest) (*responseModel.UpdateProductResponse, error) {

	if req.Name != "" && strings.Trim(req.Name, " ") == "" {
		return nil, InvalidProductNameError
	}

	if len(req.CategoryIDs) > 0 {
		uniqueCategoryIDs := helper.UniqueStringArray(req.CategoryIDs)

		if len(uniqueCategoryIDs) != len(req.CategoryIDs) {
			return nil, DuplicateCategoryError
		}

		for _, categoryID := range req.CategoryIDs {
			if _, err := pgModel.UUIDFromString(categoryID); err != nil {
				return nil, InvalidCategoryIDError
			}
		}
	}

	if req.Price < 0 {
		return nil, InvalidProductPriceError
	}

	if req.Slug != "" {
		if strings.Trim(req.Slug, " ") == "" {
			return nil, ProductSlugIsRequiredError
		}

		slugRegex, _ := regexp.Compile(regex.SlugRegex)

		if !slugRegex.MatchString(req.Slug) {
			return nil, InvalidProductSlugError
		}
	}

	if len(req.SubPhotos) > 3 {
		return nil, InvalidSecondaryPhotoQuantityError
	}

	return mw.Service.UpdateProduct(ctx, req)
}
