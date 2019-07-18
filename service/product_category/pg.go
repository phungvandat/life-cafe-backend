package productcategory

import (
	"context"

	"github.com/jinzhu/gorm"
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
)

// pgService implmenter for User serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService new pg service
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func (s *pgService) Create(ctx context.Context, req requestModel.CreateProductCategoryRequest) (*responseModel.CreateProductCategoryResponse, error) {
	productCategory := &pgModel.ProductCategory{
		Slug:  req.Slug,
		Name:  req.Name,
		Photo: req.Photo,
	}
	var err error
	if req.ParentCategoryID != "" {
		parentCategoryID, _ := pgModel.UUIDFromString(req.ParentCategoryID)
		parentCategory := &pgModel.ProductCategory{
			Model: pgModel.Model{
				ID: parentCategoryID,
			},
		}
		err = s.db.Find(parentCategory, parentCategory).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, ParentCategoryIsNotExistError
			}
			return nil, err
		}
		productCategory.ParentCategoryID = &parentCategoryID
	}

	productCategorySlug := &pgModel.ProductCategory{
		Slug: req.Slug,
	}

	err = s.db.Find(productCategorySlug, productCategorySlug).Error

	if err == nil {
		return nil, ProductCategorySlugAlreadyExistError
	}

	err = s.db.Create(productCategory).Error

	if err != nil {
		return nil, err
	}

	return &responseModel.CreateProductCategoryResponse{
		ProductCategory: productCategory,
	}, nil
}

func (s *pgService) GetProductCategory(ctx context.Context, req requestModel.GetProductCategoryRequest) (*responseModel.GetProductCategoryResponse, error) {
	productCategoryIDUUID, _ := pgModel.UUIDFromString(req.ParamProductCategoryID)
	productCategory := &pgModel.ProductCategory{
		Model: pgModel.Model{
			ID: productCategoryIDUUID,
		},
	}

	err := s.db.Find(productCategory, productCategory).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ProductCategoryNotExistError
		}
		return nil, err
	}

	productCategoryRes := &responseModel.GetProductCategoryResponse{
		ProductCategory: &responseModel.ProductCategory{
			ProductCategory: productCategory,
		},
	}

	if productCategory.ParentCategoryID != nil {
		parentCategory := &pgModel.ProductCategory{
			Model: pgModel.Model{
				ID: *productCategory.ParentCategoryID,
			},
		}
		err = s.db.Find(parentCategory, parentCategory).Error
		if err == nil {
			productCategoryRes.ProductCategory.ParentCategory = parentCategory
		}
	}

	return productCategoryRes, nil
}

func (s *pgService) GetProductCategories(ctx context.Context, req requestModel.GetProductCategoriesRequest) (*responseModel.GetProductCategoriesResponse, error) {
	skip := req.Skip
	limit := req.Limit
	if req.Skip == "" {
		skip = "-1"
	}

	if req.Limit == "" {
		limit = "-1"
	}

	arrProductCategory := []struct {
		*pgModel.ProductCategory
		ParentName  *string `json:"parent_name"`
		ParentColor *string `json:"parent_color"`
	}{}
	err := s.db.Limit(limit).Offset(skip).Table("product_categories").
		Select("product_categories.*, parent.name as parent_name, parent.color as parent_color").
		Joins("LEFT JOIN product_categories as parent ON product_categories.parent_category_id = parent.id").
		Scan(&arrProductCategory).Error

	if err != nil {
		return nil, err
	}

	listProductCategory := []*responseModel.ProductCategory{}
	for _, item := range arrProductCategory {
		productCategory := &responseModel.ProductCategory{
			ProductCategory: item.ProductCategory,
		}
		if item.ParentCategoryID != nil {
			productCategory.ParentCategory = &pgModel.ProductCategory{
				Model: pgModel.Model{
					ID: *item.ParentCategoryID,
				},
				Name:  *item.ParentName,
				Color: *item.ParentColor,
			}
		}
		listProductCategory = append(listProductCategory, productCategory)
	}

	return &responseModel.GetProductCategoriesResponse{
		ProductCategories: listProductCategory,
	}, nil
}

func (s *pgService) UpdateProductCategory(ctx context.Context, req requestModel.UpdateProductCategoryRequest) (*responseModel.UpdateProductCategoryResponse, error) {
	productCategoryRes := &responseModel.UpdateProductCategoryResponse{
		ProductCategory: &responseModel.ProductCategory{},
	}
	productCategoryID, _ := pgModel.UUIDFromString(req.ParamProductCategoryID)

	productCategory := &pgModel.ProductCategory{
		Model: pgModel.Model{
			ID: productCategoryID,
		},
	}

	err := s.db.Find(productCategory, productCategory).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ProductCategoryNotExistError
		}
		return nil, err
	}

	var parentCategory *pgModel.ProductCategory
	if req.ParentCategoryID != "" {
		parentCategoryID, _ := pgModel.UUIDFromString(req.ParentCategoryID)
		parentCategory = &pgModel.ProductCategory{
			Model: pgModel.Model{
				ID: parentCategoryID,
			},
		}

		err = s.db.Find(parentCategory, parentCategory).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = ParentCategoryIsNotExistError
			}
			return nil, err
		}
		productCategory.ParentCategoryID = &parentCategoryID
	} else if productCategory.ParentCategoryID != nil {
		parentCategory = &pgModel.ProductCategory{
			Model: pgModel.Model{
				ID: *productCategory.ParentCategoryID,
			},
		}

		err = s.db.Find(parentCategory, parentCategory).Error
		if err != nil {
			parentCategory = nil
		}
	}

	if parentCategory != nil {
		productCategoryRes.ProductCategory.ParentCategory = &pgModel.ProductCategory{
			Model: pgModel.Model{
				ID: parentCategory.ID,
			},
			Name:  parentCategory.Name,
			Color: parentCategory.Color,
		}
	}

	if req.Name != "" {
		productCategory.Name = req.Name
	}

	if req.Slug != "" {
		productCategorySlug := &pgModel.ProductCategory{}

		err = s.db.Where("id != ? AND slug = ?", productCategory.ID, req.Slug).Find(productCategorySlug).Error

		if err == nil {
			return nil, ProductCategorySlugAlreadyExistError
		}
		productCategory.Slug = req.Slug
	}

	if req.Photo != "" {
		productCategory.Photo = req.Photo
	}
	err = s.db.Save(productCategory).Error

	if err != nil {
		return nil, err
	}

	productCategoryRes.ProductCategory.ProductCategory = productCategory

	return productCategoryRes, nil
}
