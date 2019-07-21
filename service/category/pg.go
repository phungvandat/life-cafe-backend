package category

import (
	"context"
	"sync"

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
		parentCategory, err := s.getParentCategory(ctx, *category)

		if err != nil {
			return nil, err
		}
		categoryRes.Category.ParentCategory = parentCategory
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

	arrCategory := []*pgModel.Category{}

	query := ""
	if req.ParentCategoryExist == "false" {
		query += "parent_category_id IS NULL"
	} else if req.ParentCategoryExist == "true" {
		query += "parent_category_id IS NOT NULL"
	}

	if req.Slug != "" {
		query += " AND slug ='" + req.Slug + "'"
	}

	err := s.db.Limit(limit).Offset(skip).Table("categories").
		Where(query).
		Scan(&arrCategory).Error

	if err != nil {
		return nil, err
	}

	listCategory := []*responseModel.Category{}
	for _, item := range arrCategory {
		childrens, err := s.getChildrenCategories(ctx, *item)
		if err != nil {
			return nil, err
		}
		category := &responseModel.Category{
			Category:           item,
			ChildrenCategories: childrens,
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
				err = errors.ParentCategoryIsNotExistError
			}
			return nil, err
		}
		category.ParentCategoryID = &parentCategoryID
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

func (s *pgService) getParentCategory(ctx context.Context, category pgModel.Category) (*responseModel.Category, error) {
	res := &responseModel.Category{}

	categoryRes := &pgModel.Category{
		Model: pgModel.Model{
			ID: *category.ParentCategoryID,
		},
	}

	err := s.db.Find(categoryRes, categoryRes).Error

	if err != nil {
		return res, err
	}

	res.Category = &pgModel.Category{
		Model: pgModel.Model{
			ID: categoryRes.ID,
		},
		Name:  categoryRes.Name,
		Slug:  categoryRes.Slug,
		Color: categoryRes.Color,
	}

	if categoryRes.ParentCategoryID != nil {
		parentCategory, err := s.getParentCategory(ctx, *categoryRes)
		if err != nil {
			return res, err
		}
		res.ParentCategory = parentCategory
	}

	return res, nil
}

func (s *pgService) getChildrenCategories(ctx context.Context, category pgModel.Category) ([]*responseModel.Category, error) {
	res := []*responseModel.Category{}

	childrenCategories := []*pgModel.Category{}

	err := s.db.Table("categories").
		Where("parent_category_id = ?", category.ID).
		Scan(&childrenCategories).Error

	if err != nil {
		return res, err
	}
	wg := sync.WaitGroup{}

	for _, item := range childrenCategories {
		wg.Add(1)
		go func(item *pgModel.Category) {
			defer wg.Done()
			childrens, errFunc := s.getChildrenCategories(ctx, *item)
			if errFunc != nil {
				err = errFunc
				return
			}
			categoryRes := &responseModel.Category{
				Category: &pgModel.Category{
					Model: pgModel.Model{
						ID: item.ID,
					},
					Name:  item.Name,
					Slug:  item.Slug,
					Color: item.Color,
				},
				ChildrenCategories: childrens,
			}
			res = append(res, categoryRes)
		}(item)
	}
	wg.Wait()
	return res, nil
}
