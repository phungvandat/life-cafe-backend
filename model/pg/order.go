package pg

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Order struct
type Order struct {
	Model
	Type                string `json:"type,omitempty"`
	CreatorID           *UUID  `json:"creatorIDd,omitempty"`
	Note                string `json:"note,omitempty"`
	CustomerID          *UUID  `json:"customerIDd,omitempty"`
	Status              string `json:"status,omitempty"`
	ImplementerID       *UUID  `json:"implementerIDd,omitempty"`
	ReceiverPhoneNumber string `json:"receiverPhoneNumber,omitempty"`
	ReceiverAddress     string `json:"receiverAddress,omitempty"`
	ReceiverFullname    string `json:"receiverFullname,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *Order) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// ProductOrder struct
type ProductOrder struct {
	Model
	ProductID     *UUID `json:"productIDd,omitempty"`
	OrderID       *UUID `json:"orderIDd,omitempty"`
	OrderQuantity int   `json:"orderQuantity,omitempty"`
	RealPrice     int   `json:"realPrice,omitempty"`
}

// BeforeCreate prepare data before create data
func (u *ProductOrder) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
