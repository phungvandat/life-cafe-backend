package pg

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// ProductOrder struct
type ProductOrder struct {
	Model
	Type                string `json:"type,omitempty"`
	CreatorID           string `json:"creator_id,omitempty"`
	Note                string `json:"note,omitempty"`
	CustomerID          string `json:"customer_id,omitempty"`
	Status              string `json:"status,omitempty"`
	ImplementerID       string `json:"implementer_id,omitempty"`
	ReceiverPhoneNumber string `json:"receiver_phone_number,omitempty"`
	ReceiverAddress     string `json:"receiver_address,omitempty"`
	ReceiverFullname    string `json:"receiver_fullname,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *ProductOrder) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
