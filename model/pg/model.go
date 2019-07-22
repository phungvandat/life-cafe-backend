package pg

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Model is base model
type Model struct {
	ID        UUID       `sql:",type:uuid" json:"id,omitempty"`
	CreatedAt *time.Time `sql:"default:now()" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty,omitempty"`
}

// BeforeCreate prepare data before create data
func (at *Model) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.NewV4())
	if err != nil {
		return err
	}
	err = scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		return err
	}
	err = scope.SetColumn("UpdatedAt", time.Now())
	if err != nil {
		return err
	}
	return nil
}
