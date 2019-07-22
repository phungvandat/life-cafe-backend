package pg

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/phungvandat/life-cafe-backend/util/helper"
	uuid "github.com/satori/go.uuid"
)

// Category struct save information of an category
type Category struct {
	Model
	Name             string `json:"name,omitempty"`
	Photo            string `json:"photo,omitempty"`
	ParentCategoryID *UUID  `json:"parentCategoryID,omitempty"`
	Slug             string `json:"slug,omitempty"`
	Color            string `json:"color,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *Category) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("Color", helper.GetRandomColorInHex())
	scope.SetColumn("Slug", strings.ToLower(u.Slug))
	return nil
}

// BeforeSave prepare data before save data
func (u *Category) BeforeSave(scope *gorm.Scope) error {
	if u.Slug != "" {
		scope.SetColumn("Slug", strings.ToLower(u.Slug))
	}
	return nil
}
