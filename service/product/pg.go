package product

import (
	"context"
	"sync"

	"github.com/jinzhu/gorm"
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	productCategorySvc "github.com/phungvandat/life-cafe-backend/service/product_category"
	"github.com/phungvandat/life-cafe-backend/util/helper"
)

// pgService implmenter for auth service in postgres
type pgService struct {
	db                 *gorm.DB
	productCategorySvc productCategorySvc.Service
	spRollback         helper.SagasService
}

// NewPGService new pg service
func NewPGService(db *gorm.DB, productCategorySvc productCategorySvc.Service, spRollback helper.SagasService) Service {
	return &pgService{
		db:                 db,
		productCategorySvc: productCategorySvc,
		spRollback:         spRollback,
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
		return res, ProductSlugExistError
	}

	// product category
	categories := []*pgModel.ProductCategory{}
	for _, categoryID := range req.CategoryIDs {
		categoryIDUUID, _ := pgModel.UUIDFromString(categoryID)
		category := &pgModel.ProductCategory{
			Model: pgModel.Model{
				ID: categoryIDUUID,
			},
		}
		err = s.db.Find(category, category).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return res, ProductCategoryNotExistError
			}
			return res, err
		}
		categories = append(categories, &pgModel.ProductCategory{
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
		productCategory := &pgModel.Productcategory{
			ProductID:         &product.ID,
			ProductCategoryID: &categoryIDUUID,
		}
		err = tx.Create(productCategory).Error
		if err != nil {
			return res, err
		}
	}

	res.Product.Product = product

	s.removeProductSubPhotos(ctx, product)

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
			err = ProductNotExistError
		}
		return res, err
	}

	// Get product category
	categories, err := s.getProductCategories(ctx, (product.ID).String())

	if err != nil {
		return res, err
	}

	res.Product.Categories = categories

	// Sub photos
	subPhotos := s.getProductSubPhotos(ctx, product)
	s.removeProductSubPhotos(ctx, product)
	res.Product.SubPhotos = subPhotos
	res.Product.Product = product

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
		subPhotos := s.getProductSubPhotos(ctx, &product)
		s.removeProductSubPhotos(ctx, &product)
		productRes.SubPhotos = subPhotos
		productRes.Product = &product

		wg.Add(1)
		go func(productID string) {
			defer wg.Done()
			categories, err := s.getProductCategories(ctx, productID)
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
			err = ProductNotExistError
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
		s.removeProductSubPhotos(ctx, product)
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
			return res, ProductSlugExistError
		}
		product.Slug = req.Slug
	}

	if req.Barcode != "" && req.Barcode != product.Barcode {
		product.Barcode = req.Barcode
	}

	if req.Description != "" && req.Description != product.Description {
		product.Description = req.Description
	}

	preCategoryIDs := []string{}

	preCategoryIDsStruct := []struct {
		ProductCategoryID string `json:"product_category_id,omitempty"`
	}{}

	err = s.db.Model(&pgModel.Productcategory{}).Where("product_id = ?", product.ID).Select("product_category_id").Scan(&preCategoryIDsStruct).Error

	if err != nil {
		return res, err
	}

	for _, item := range preCategoryIDsStruct {
		preCategoryIDs = append(preCategoryIDs, item.ProductCategoryID)
	}

	categoryIDs := []string{}
	if len(req.CategoryIDs) > 0 {
		sameCategoryIDs := helper.GetSameElementInArrays(preCategoryIDs, req.CategoryIDs)
		deleteCategoryIDs := helper.DifferenceArray(sameCategoryIDs, preCategoryIDs)
		createCategoryIDs := helper.DifferenceArray(sameCategoryIDs, req.CategoryIDs)
		if len(deleteCategoryIDs) > 0 {
			deleteQuery := "(product_id = '" + req.ParamProductID + "' AND product_category_id = '" + deleteCategoryIDs[0] + "')"
			for _, deleteID := range deleteCategoryIDs[1:] {
				deleteQuery += "OR (product_id = '" + req.ParamProductID + "' AND product_category_id = '" + deleteID + "')"
			}
			err = tx.Unscoped().Where(deleteQuery).Delete(&pgModel.Productcategory{}).Error
			if err != nil {
				return res, err
			}
		}

		if len(createCategoryIDs) > 0 {
			for _, createID := range createCategoryIDs {
				productCategoryIDUUID, _ := pgModel.UUIDFromString(createID)
				productcategory := &pgModel.Productcategory{
					ProductID:         &product.ID,
					ProductCategoryID: &productCategoryIDUUID,
				}
				err = tx.Create(productcategory).Error
				if err != nil {
					return res, err
				}
			}
		}
		categoryIDs = append(sameCategoryIDs, createCategoryIDs...)
	}

	err = tx.Save(product).Error

	if err != nil {
		return res, err
	}

	categories, err := s.getCategories(ctx, categoryIDs)

	if err != nil {
		return res, err
	}

	res.Product.Product = product
	res.Product.Categories = categories
	subPhotos := s.getProductSubPhotos(ctx, product)
	res.Product.SubPhotos = subPhotos
	s.removeProductSubPhotos(ctx, product)

	return res, nil
}

func (s *pgService) getCategories(ctx context.Context, categoryIDs []string) ([]*pgModel.ProductCategory, error) {
	categories := []*pgModel.ProductCategory{}

	wg := sync.WaitGroup{}
	for _, categoryID := range categoryIDs {
		wg.Add(1)
		go func(categoryID string) {
			defer wg.Done()
			category, err := s.productCategorySvc.GetProductCategory(ctx,
				requestModel.GetProductCategoryRequest{
					ParamProductCategoryID: categoryID,
				})

			if err == nil {
				categories = append(categories, &pgModel.ProductCategory{
					Model: pgModel.Model{
						ID: category.ProductCategory.ID,
					},
					Name:  category.ProductCategory.Name,
					Color: category.ProductCategory.Color,
				})
			}
		}(categoryID)
	}
	wg.Wait()

	return categories, nil
}

func (s *pgService) getProductCategories(ctx context.Context, productID string) ([]*pgModel.ProductCategory, error) {
	categories := []*pgModel.ProductCategory{}
	productCategories := []pgModel.Productcategory{}

	err := s.db.Where("product_id = ?", productID).Find(&productCategories).Error
	if err != nil {
		return categories, err
	}

	wg := sync.WaitGroup{}
	for _, pc := range productCategories {
		wg.Add(1)
		go func(categoryID string) {
			defer wg.Done()
			category, err := s.productCategorySvc.GetProductCategory(ctx,
				requestModel.GetProductCategoryRequest{
					ParamProductCategoryID: categoryID,
				})

			if err == nil {
				categories = append(categories, &pgModel.ProductCategory{
					Model: pgModel.Model{
						ID: category.ProductCategory.ID,
					},
					Name:  category.ProductCategory.Name,
					Color: category.ProductCategory.Color,
				})
			}
		}((pc.ProductCategoryID).String())
	}
	wg.Wait()

	return categories, nil
}

func (s *pgService) getProductcategogyStringIDArray(ctx context.Context, productCategories []pgModel.Productcategory) []string {
	stringIDArray := []string{}

	for _, pc := range productCategories {
		stringIDArray = append(stringIDArray, (pc.ProductCategoryID).String())
	}
	return stringIDArray
}

func (s *pgService) getProductSubPhotos(ctx context.Context, product *pgModel.Product) []string {
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

func (s *pgService) removeProductSubPhotos(ctx context.Context, product *pgModel.Product) {
	product.FirstSubPhoto = ""
	product.SecondSubPhoto = ""
	product.ThirdSubPhoto = ""
}
