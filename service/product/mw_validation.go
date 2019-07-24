package product

import (
	"context"
	"strconv"
	"strings"

	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
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
		return nil, errors.ProductNameIsRequiredError
	}

	if req.MainPhoto == "" {
		return nil, errors.ProductMainPhotoIsRequiredError
	}

	if req.Quantity < 0 {
		return nil, errors.InvalidProductQuantityError
	}

	if req.CategoryIDs == nil || len(req.CategoryIDs) <= 0 {
		return nil, errors.CategoryOfProductIsRequiredError
	}

	uniqueCategoryIDs := helper.UniqueStringArray(req.CategoryIDs)

	if len(req.CategoryIDs) != len(uniqueCategoryIDs) {
		return nil, errors.DuplicateCategoryError
	}

	for _, categoryID := range req.CategoryIDs {
		if _, err := pgModel.UUIDFromString(categoryID); err != nil {
			return nil, errors.InvalidCategoryIDError
		}
	}

	if req.Price < 0 {
		return nil, errors.InvalidProductPriceError
	}

	if strings.Trim(req.Slug, " ") == "" {
		return nil, errors.ProductSlugIsRequiredError
	}

	if !regex.IsSlugValid(req.Slug) {
		return nil, errors.InvalidProductSlugError
	}

	if len(req.SubPhotos) > 3 {
		return nil, errors.InvalidSecondaryPhotoQuantityError
	}

	for _, photo := range req.SubPhotos {
		if photo.URL == "" {
			return nil, errors.InvalidPhotoURLError
		}
		if _, err := pgModel.UUIDFromString(photo.ID); photo.ID != "" && err != nil {
			return nil, errors.InvalidPhotoIDError
		}
	}

	return mw.Service.CreateProduct(ctx, req)
}

func (mw validationMiddleware) GetProduct(ctx context.Context, req requestModel.GetProductRequest) (*responseModel.GetProductResponse, error) {
	if _, err := pgModel.UUIDFromString(req.ParamProductID); req.ParamProductID != "" && err != nil {
		return nil, errors.InvalidProductIDTypeError
	}
	return mw.Service.GetProduct(ctx, req)
}

func (mw validationMiddleware) GetProducts(ctx context.Context, req requestModel.GetProductsRequest) (*responseModel.GetProductsResponse, error) {
	if _, err := strconv.ParseInt(req.Skip, 10, 32); req.Skip != "" && err != nil {
		return nil, errors.InvalidSkipError
	}

	if _, err := strconv.ParseInt(req.Limit, 10, 32); req.Limit != "" && err != nil {
		return nil, errors.InvalidLimitError
	}

	return mw.Service.GetProducts(ctx, req)
}

func (mw validationMiddleware) UpdateProduct(ctx context.Context, req requestModel.UpdateProductRequest) (*responseModel.UpdateProductResponse, error) {

	if req.Name != "" && strings.Trim(req.Name, " ") == "" {
		return nil, errors.InvalidProductNameError
	}

	if len(req.CategoryIDs) > 0 {
		uniqueCategoryIDs := helper.UniqueStringArray(req.CategoryIDs)

		if len(uniqueCategoryIDs) != len(req.CategoryIDs) {
			return nil, errors.DuplicateCategoryError
		}

		for _, categoryID := range req.CategoryIDs {
			if _, err := pgModel.UUIDFromString(categoryID); err != nil {
				return nil, errors.InvalidCategoryIDError
			}
		}
	}

	if req.Price < 0 {
		return nil, errors.InvalidProductPriceError
	}

	if req.Slug != "" {
		if strings.Trim(req.Slug, " ") == "" {
			return nil, errors.ProductSlugIsRequiredError
		}

		if !regex.IsSlugValid(req.Slug) {
			return nil, errors.InvalidProductSlugError
		}
	}

	if len(req.SubPhotos) > 3 {
		return nil, errors.InvalidSecondaryPhotoQuantityError
	}

	for _, photo := range req.SubPhotos {
		if photo.URL == "" {
			return nil, errors.InvalidPhotoURLError
		}
		if _, err := pgModel.UUIDFromString(photo.ID); photo.ID != "" && err != nil {
			return nil, errors.InvalidPhotoIDError
		}
	}

	if req.Quantity != nil {
		if *req.Quantity < 0 {
			return nil, errors.InvalidProductQuantityError
		}
	}

	return mw.Service.UpdateProduct(ctx, req)
}
