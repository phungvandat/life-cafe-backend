package domain

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Model is base model
type Model struct {
	ID        UUID       `sql:",type:uuid" json:"id,omitempty"`
	CreatedAt *time.Time `sql:"default:now()" json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty,omitempty"`
}

// BeforeCreate prepare data before create data
func (at *Model) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.NewV4())
	if err != nil {
		return err
	}
	fmt.Println("hihihi", uuid.NewV4())
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
