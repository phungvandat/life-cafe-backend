package product

import (
	"context"
	"sync"

	"github.com/jinzhu/gorm"
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	categorySvc "github.com/phungvandat/life-cafe-backend/service/category"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
	"github.com/phungvandat/life-cafe-backend/util/helper"
)

// pgService implmenter for auth service in postgres
type pgService struct {
	db          *gorm.DB
	categorySvc categorySvc.Service
	spRollback  helper.SagasService
}

// NewPGService new pg service
func NewPGService(db *gorm.DB, categorySvc categorySvc.Service, spRollback helper.SagasService) Service {
	return &pgService{
		db:          db,
		categorySvc: categorySvc,
		spRollback:  spRollback,
	}
}

func (s *pgService) RollbackTransaction(_ context.Context, transactionID string) error {
	return s.spRollback.RollbackTransaction(transactionID)
}

func (s *pgService) CommitTransaction(_ context.Context, transactionID string) error {
	return s.spRollback.CommitTransaction(transactionID)
}

func (s *pgService) CreateProduct(ctx context.Context, req requestModel.CreateProductRequest) (*responseModel.CreateProductResponse, error) {
	tx := s.db.Begin()
	transactionID := (pgModel.NewUUID()).String()
	s.spRollback.NewTransaction(transactionID, tx)
	res := &responseModel.CreateProductResponse{
		TransactionID: &transactionID,
		Product:       &responseModel.Product{},
	}

	product := &pgModel.Product{
		Name:        req.Name,
		MainPhoto:   req.MainPhoto,
		Description: req.Description,
		Price:       req.Price,
		Barcode:     req.Barcode,
		Slug:        req.Slug,
	}

	slugProduct := &pgModel.Product{
		Slug: req.Slug,
	}

	err := s.db.Find(slugProduct, slugProduct).Error

	if err == nil {
		return res, errors.ProductSlugExistError
	}

	// category
	categories := []*pgModel.Category{}
	for _, categoryID := range req.CategoryIDs {
		categoryIDUUID, _ := pgModel.UUIDFromString(categoryID)
		category := &pgModel.Category{
			Model: pgModel.Model{
				ID: categoryIDUUID,
			},
		}
		err = s.db.Find(category, category).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return res, errors.CategoryNotExistError
			}
			return res, err
		}
		categories = append(categories, &pgModel.Category{
			Name:  category.Name,
			Color: category.Color,
			Model: pgModel.Model{
				ID: category.ID,
			},
		})
	}
	res.Product.Categories = categories

	// product sub photos
	subPhotos := []string{}
	for index, subPhoto := range req.SubPhotos {
		if index == 0 {
			product.FirstSubPhoto = subPhoto
		} else if index == 1 {
			product.SecondSubPhoto = subPhoto
		} else if index == 2 {
			product.ThirdSubPhoto = subPhoto
		}
		subPhotos = append(subPhotos, subPhoto)
	}
	res.Product.SubPhotos = subPhotos

	err = tx.Create(product).Error
	if err != nil {
		return res, err
	}

	for _, categoryID := range req.CategoryIDs {
		categoryIDUUID, _ := pgModel.UUIDFromString(categoryID)
		category := &pgModel.ProductCategory{
			ProductID:  &product.ID,
			CategoryID: &categoryIDUUID,
		}
		err = tx.Create(category).Error
		if err != nil {
			return res, err
		}
	}

	res.Product.Product = s.removeProductSubPhotos(ctx, *product)

	return res, nil
}

func (s *pgService) GetProduct(ctx context.Context, req requestModel.GetProductRequest) (*responseModel.GetProductResponse, error) {
	res := &responseModel.GetProductResponse{
		Product: &responseModel.Product{},
	}

	productIDUUID, _ := pgModel.UUIDFromString(req.ParamProductID)

	product := &pgModel.Product{
		Model: pgModel.Model{
			ID: productIDUUID,
		},
	}

	err := s.db.Find(product, product).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.ProductNotExistError
		}
		return res, err
	}

	// Get category
	categories, err := s.getCategoriesForProduct(ctx, (product.ID).String())

	if err != nil {
		return res, err
	}

	res.Product.Categories = categories

	// Sub photos
	subPhotos := s.getProductSubPhotos(ctx, *product)
	res.Product.SubPhotos = subPhotos
	res.Product.Product = s.removeProductSubPhotos(ctx, *product)

	return res, nil
}

func (s *pgService) GetProducts(ctx context.Context, req requestModel.GetProductsRequest) (*responseModel.GetProductsResponse, error) {
	res := &responseModel.GetProductsResponse{}
	skip := req.Skip
	limit := req.Limit
	if req.Skip == "" {
		skip = "-1"
	}

	if req.Limit == "" {
		limit = "-1"
	}

	products := []pgModel.Product{}

	err := s.db.Limit(limit).Offset(skip).Find(&products).Error

	if err != nil {
		return res, err
	}

	productsRes := []*responseModel.Product{}

	wg := sync.WaitGroup{}
	for _, product := range products {
		productRes := &responseModel.Product{}
		subPhotos := s.getProductSubPhotos(ctx, product)
		productRes.SubPhotos = subPhotos
		productRes.Product = s.removeProductSubPhotos(ctx, product)

		wg.Add(1)
		go func(productID string) {
			defer wg.Done()
			categories, err := s.getCategoriesForProduct(ctx, productID)
			if err == nil {
				productRes.Categories = categories
			}
			productsRes = append(productsRes, productRes)
		}((product.ID).String())
	}
	wg.Wait()

	res.Products = productsRes
	return res, nil
}

func (s *pgService) UpdateProduct(ctx context.Context, req requestModel.UpdateProductRequest) (*responseModel.UpdateProductResponse, error) {
	tx := s.db.Begin()
	transactionID := (pgModel.NewUUID()).String()
	s.spRollback.NewTransaction(transactionID, tx)

	res := &responseModel.UpdateProductResponse{
		Product:       &responseModel.Product{},
		TransactionID: &transactionID,
	}

	productID, _ := pgModel.UUIDFromString(req.ParamProductID)

	product := &pgModel.Product{
		Model: pgModel.Model{
			ID: productID,
		},
	}

	err := s.db.Find(product, product).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.ProductNotExistError
		}
		return res, err
	}

	if req.Name != "" && req.Name != product.Name {
		product.Name = req.Name
	}

	if req.MainPhoto != "" && req.MainPhoto != product.MainPhoto {
		product.MainPhoto = req.MainPhoto
	}

	// product sub photos
	if len(req.SubPhotos) > 0 {
		for index, subPhoto := range req.SubPhotos {
			if index == 0 {
				product.FirstSubPhoto = subPhoto
			} else if index == 1 {
				product.SecondSubPhoto = subPhoto
			} else if index == 2 {
				product.ThirdSubPhoto = subPhoto
			}
		}
	}

	if req.Price >= 0 && req.Price != product.Price {
		product.Price = req.Price
	}

	if req.Flag != 0 && req.Flag != product.Flag {
		product.Flag = req.Flag
	}

	if req.Slug != "" && req.Slug != product.Slug {
		slugProduct := &pgModel.Product{
			Slug: req.Slug,
		}

		err := s.db.Find(slugProduct, slugProduct).Error

		if err == nil {
			return res, errors.ProductSlugExistError
		}
		product.Slug = req.Slug
	}

	if req.Barcode != "" && req.Barcode != product.Barcode {
		product.Barcode = req.Barcode
	}

	if req.Description != "" && req.Description != product.Description {
		product.Description = req.Description
	}

	if req.Quantity != nil && *req.Quantity != product.Quantity {
		product.Quantity = *req.Quantity
	}

	preCategoryIDs := []string{}

	preCategoryIDsStruct := []struct {
		CategoryID string `json:"category_id,omitempty"`
	}{}

	err = s.db.Model(&pgModel.ProductCategory{}).Where("product_id = ?", product.ID).Select("category_id").Scan(&preCategoryIDsStruct).Error

	if err != nil {
		return res, err
	}

	for _, item := range preCategoryIDsStruct {
		preCategoryIDs = append(preCategoryIDs, item.CategoryID)
	}

	categoryIDs := []string{}
	if len(req.CategoryIDs) > 0 {
		sameCategoryIDs := helper.GetSameElementInArrays(preCategoryIDs, req.CategoryIDs)
		deleteCategoryIDs := helper.DifferenceArray(sameCategoryIDs, preCategoryIDs)
		createCategoryIDs := helper.DifferenceArray(sameCategoryIDs, req.CategoryIDs)
		if len(deleteCategoryIDs) > 0 {
			deleteQuery := "(product_id = '" + req.ParamProductID + "' AND category_id = '" + deleteCategoryIDs[0] + "')"
			for _, deleteID := range deleteCategoryIDs[1:] {
				deleteQuery += "OR (product_id = '" + req.ParamProductID + "' AND category_id = '" + deleteID + "')"
			}
			err = tx.Unscoped().Where(deleteQuery).Delete(&pgModel.ProductCategory{}).Error
			if err != nil {
				return res, err
			}
		}

		if len(createCategoryIDs) > 0 {
			for _, createID := range createCategoryIDs {
				categoryIDUUID, _ := pgModel.UUIDFromString(createID)
				productcategory := &pgModel.ProductCategory{
					ProductID:  &product.ID,
					CategoryID: &categoryIDUUID,
				}
				err = tx.Create(productcategory).Error
				if err != nil {
					return res, err
				}
			}
		}
		categoryIDs = append(sameCategoryIDs, createCategoryIDs...)
	} else {
		categoryIDs = preCategoryIDs
	}

	err = tx.Save(product).Error

	if err != nil {
		return res, err
	}

	categories, err := s.getCategoriesByIDs(ctx, categoryIDs)

	if err != nil {
		return res, err
	}

	res.Product.Categories = categories
	subPhotos := s.getProductSubPhotos(ctx, *product)
	res.Product.SubPhotos = subPhotos
	res.Product.Product = s.removeProductSubPhotos(ctx, *product)

	return res, nil
}

func (s *pgService) getCategoriesByIDs(ctx context.Context, categoryIDs []string) ([]*pgModel.Category, error) {
	categories := []*pgModel.Category{}

	wg := sync.WaitGroup{}
	for _, categoryID := range categoryIDs {
		wg.Add(1)
		go func(categoryID string) {
			defer wg.Done()
			category, err := s.categorySvc.GetCategory(ctx,
				requestModel.GetCategoryRequest{
					ParamCategoryID: categoryID,
				})

			if err == nil {
				categories = append(categories, &pgModel.Category{
					Model: pgModel.Model{
						ID: category.Category.ID,
					},
					Name:  category.Category.Name,
					Color: category.Category.Color,
				})
			}
		}(categoryID)
	}
	wg.Wait()

	return categories, nil
}

func (s *pgService) getCategoriesForProduct(ctx context.Context, productID string) ([]*pgModel.Category, error) {
	categories := []*pgModel.Category{}
	productCategories := []pgModel.ProductCategory{}

	err := s.db.Where("product_id = ?", productID).Find(&productCategories).Error
	if err != nil {
		return categories, err
	}

	wg := sync.WaitGroup{}
	for _, pc := range productCategories {
		wg.Add(1)
		go func(categoryID string) {
			defer wg.Done()
			category, err := s.categorySvc.GetCategory(ctx,
				requestModel.GetCategoryRequest{
					ParamCategoryID: categoryID,
				})

			if err == nil {
				categories = append(categories, &pgModel.Category{
					Model: pgModel.Model{
						ID: category.Category.ID,
					},
					Name:  category.Category.Name,
					Color: category.Category.Color,
				})
			}
		}((pc.CategoryID).String())
	}
	wg.Wait()

	return categories, nil
}

func (s *pgService) getProductcategogyStringIDArray(ctx context.Context, productCategories []pgModel.ProductCategory) []string {
	stringIDArray := []string{}

	for _, pc := range productCategories {
		stringIDArray = append(stringIDArray, (pc.CategoryID).String())
	}
	return stringIDArray
}

func (s *pgService) getProductSubPhotos(ctx context.Context, product pgModel.Product) []string {
	subPhotos := []string{}

	if product.FirstSubPhoto != "" {
		subPhotos = append(subPhotos, product.FirstSubPhoto)
	}

	if product.SecondSubPhoto != "" {
		subPhotos = append(subPhotos, product.SecondSubPhoto)
	}

	if product.ThirdSubPhoto != "" {
		subPhotos = append(subPhotos, product.ThirdSubPhoto)
	}

	return subPhotos
}

func (s *pgService) removeProductSubPhotos(ctx context.Context, product pgModel.Product) *pgModel.Product {
	product.FirstSubPhoto = ""
	product.SecondSubPhoto = ""
	product.ThirdSubPhoto = ""
	return &product
}
