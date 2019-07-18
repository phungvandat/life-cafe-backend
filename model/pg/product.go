package pg

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/phungvandat/life-cafe-backend/util/helper"
	uuid "github.com/satori/go.uuid"
)

// Product struct
type Product struct {
	Model
	Name           string `json:"name,omitempty"`
	MainPhoto      string `json:"main_photo,omitempty"`
	FirstSubPhoto  string `json:"first_sub_photo,omitempty"`
	SecondSubPhoto string `json:"second_sub_photo,omitempty"`
	ThirdSubPhoto  string `json:"third_sub_photo,omitempty"`
	Quantity       int    `json:"quantity,omitempty"`
	Price          int    `json:"price,omitempty"`
	Flag           int    `json:"flag,omitempty"`
	Color          string `json:"color,omitempty"`
	Slug           string `json:"slug,omitempty"`
	Barcode        string `json:"barcode,omitempty"`
	Description    string `json:"description,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *Product) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("Color", helper.GetRandomColorInHex())
	scope.SetColumn("Slug", strings.ToLower(u.Slug))
	return nil
}

// BeforeSave prepare data before save data
func (u *Product) BeforeSave(scope *gorm.Scope) error {
	if u.Slug != "" {
		scope.SetColumn("Slug", strings.ToLower(u.Slug))
	}
	return nil
}

// ProductCategory struct
type ProductCategory struct {
	Model
	ProductID  *UUID `json:"product_id,omitempty"`
	CategoryID *UUID `json:"category_id,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *ProductCategory) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
