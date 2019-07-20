package category

import (
	"context"

	"github.com/jinzhu/gorm"
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
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

func (s *pgService) Create(ctx context.Context, req requestModel.CreateCategoryRequest) (*responseModel.CreateCategoryResponse, error) {
	category := &pgModel.Category{
		Slug:  req.Slug,
		Name:  req.Name,
		Photo: req.Photo,
	}
	var err error
	if req.ParentCategoryID != "" {
		parentCategoryID, _ := pgModel.UUIDFromString(req.ParentCategoryID)
		parentCategory := &pgModel.Category{
			Model: pgModel.Model{
				ID: parentCategoryID,
			},
		}
		err = s.db.Find(parentCategory, parentCategory).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.ParentCategoryIsNotExistError
			}
			return nil, err
		}
		category.ParentCategoryID = &parentCategoryID
	}

	categorySlug := &pgModel.Category{
		Slug: req.Slug,
	}

	err = s.db.Find(categorySlug, categorySlug).Error

	if err == nil {
		return nil, errors.CategorySlugAlreadyExistError
	}

	err = s.db.Create(category).Error

	if err != nil {
		return nil, err
	}

	return &responseModel.CreateCategoryResponse{
		Category: category,
	}, nil
}

func (s *pgService) GetCategory(ctx context.Context, req requestModel.GetCategoryRequest) (*responseModel.GetCategoryResponse, error) {
	categoryIDUUID, _ := pgModel.UUIDFromString(req.ParamCategoryID)
	category := &pgModel.Category{
		Model: pgModel.Model{
			ID: categoryIDUUID,
		},
	}

	err := s.db.Find(category, category).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.CategoryNotExistError
		}
		return nil, err
	}

	categoryRes := &responseModel.GetCategoryResponse{
		Category: &responseModel.Category{
			Category: category,
		},
	}

	if category.ParentCategoryID != nil {
		parentCategory := &pgModel.Category{
			Model: pgModel.Model{
				ID: *category.ParentCategoryID,
			},
		}
		err = s.db.Find(parentCategory, parentCategory).Error
		if err == nil {
			categoryRes.Category.ParentCategory = &pgModel.Category{
				Model: pgModel.Model{
					ID: parentCategory.ID,
				},
				Name:  parentCategory.Name,
				Color: parentCategory.Color,
			}
		}
	}

	return categoryRes, nil
}

func (s *pgService) GetCategories(ctx context.Context, req requestModel.GetCategoriesRequest) (*responseModel.GetCategoriesResponse, error) {
	skip := req.Skip
	limit := req.Limit
	if req.Skip == "" {
		skip = "-1"
	}

	if req.Limit == "" {
		limit = "-1"
	}

	arrCategory := []struct {
		*pgModel.Category
		ParentName  *string `json:"parent_name"`
		ParentColor *string `json:"parent_color"`
	}{}
	err := s.db.Limit(limit).Offset(skip).Table("categories").
		Select("categories.*, parent.name as parent_name, parent.color as parent_color").
		Joins("LEFT JOIN categories as parent ON categories.parent_category_id = parent.id").
		Scan(&arrCategory).Error

	if err != nil {
		return nil, err
	}

	listCategory := []*responseModel.Category{}
	for _, item := range arrCategory {
		category := &responseModel.Category{
			Category: item.Category,
		}
		if item.ParentCategoryID != nil {
			category.ParentCategory = &pgModel.Category{
				Model: pgModel.Model{
					ID: *item.ParentCategoryID,
				},
				Name:  *item.ParentName,
				Color: *item.ParentColor,
			}
		}
		listCategory = append(listCategory, category)
	}

	return &responseModel.GetCategoriesResponse{
		Categories: listCategory,
	}, nil
}

func (s *pgService) UpdateCategory(ctx context.Context, req requestModel.UpdateCategoryRequest) (*responseModel.UpdateCategoryResponse, error) {
	categoryRes := &responseModel.UpdateCategoryResponse{
		Category: &responseModel.Category{},
	}
	categoryID, _ := pgModel.UUIDFromString(req.ParamCategoryID)

	category := &pgModel.Category{
		Model: pgModel.Model{
			ID: categoryID,
		},
	}

	err := s.db.Find(category, category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.CategoryNotExistError
		}
		return nil, err
	}

	var parentCategory *pgModel.Category
	if req.ParentCategoryID != "" {
		parentCategoryID, _ := pgModel.UUIDFromString(req.ParentCategoryID)
		parentCategory = &pgModel.Category{
			Model: pgModel.Model{
				ID: parentCategoryID,
			},
		}

		err = s.db.Find(parentCategory, parentCategory).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = errors.ParentCategoryIsNotExistError
			}
			return nil, err
		}
		category.ParentCategoryID = &parentCategoryID
	} else if category.ParentCategoryID != nil {
		parentCategory = &pgModel.Category{
			Model: pgModel.Model{
				ID: *category.ParentCategoryID,
			},
		}

		err = s.db.Find(parentCategory, parentCategory).Error
		if err != nil {
			parentCategory = nil
		}
	}

	if parentCategory != nil {
		categoryRes.Category.ParentCategory = &pgModel.Category{
			Model: pgModel.Model{
				ID: parentCategory.ID,
			},
			Name:  parentCategory.Name,
			Color: parentCategory.Color,
		}
	}

	if req.Name != "" {
		category.Name = req.Name
	}

	if req.Slug != "" {
		categorySlug := &pgModel.Category{}

		err = s.db.Where("id != ? AND slug = ?", category.ID, req.Slug).Find(categorySlug).Error

		if err == nil {
			return nil, errors.CategorySlugAlreadyExistError
		}
		category.Slug = req.Slug
	}

	if req.Photo != "" {
		category.Photo = req.Photo
	}
	err = s.db.Save(category).Error

	if err != nil {
		return nil, err
	}

	categoryRes.Category.Category = category

	return categoryRes, nil
}
