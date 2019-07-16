package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// ProductCategory struct save information of an product category
type ProductCategory struct {
	Model
	Name             string `json:"name,omitempty"`
	Photo            string `json:"photo,omitempty"`
	ParentCategoryID *UUID  `json:"parent_category_id,omitempty"`
	Slug             string `json:"slug,omitempty"`
	Color            string `json:"color,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *ProductCategory) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
