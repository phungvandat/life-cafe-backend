package pg

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Order struct
type Order struct {
	Model
	Type                string `json:"type,omitempty"`
	CreatorID           *UUID  `json:"creator_id,omitempty"`
	Note                string `json:"note,omitempty"`
	CustomerID          *UUID  `json:"customer_id,omitempty"`
	Status              string `json:"status,omitempty"`
	ImplementerID       *UUID  `json:"implementer_id,omitempty"`
	ReceiverPhoneNumber string `json:"receiver_phone_number,omitempty"`
	ReceiverAddress     string `json:"receiver_address,omitempty"`
	ReceiverFullname    string `json:"receiver_fullname,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *Order) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// ProductOrder struct
type ProductOrder struct {
	Model
	ProductID     *UUID `json:"product_id,omitempty"`
	OrderID       *UUID `json:"order_id,omitempty"`
	OrderQuantity int   `json:"order_quantity,omitempty"`
	RealPrice     int   `json:"real_price,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *ProductOrder) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
