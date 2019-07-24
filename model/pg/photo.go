package pg

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Photo struct
type Photo struct {
	Model
	URL       string `json:"url,omitempty"`
	ProductID *UUID  `json:"productID,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *Photo) BeforeCreate(scope *gorm.Scope) error {
	if u.ID.String() == "00000000-0000-0000-0000-000000000000" {
		scope.SetColumn("ID", uuid.NewV4())
	}
	return nil
}
