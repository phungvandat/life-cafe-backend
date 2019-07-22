package response

import (
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
)

// CreateOrderResponse struct
type CreateOrderResponse struct {
	Order         *Order  `json:"order,omitempty"`
	TransactionID *string `json:"transactionID,omitempty"`
}

// Order struct
type Order struct {
	*pgModel.Order
	Customer         *pgModel.User       `json:"customer,omitempty"`
	OrderProductInfo []*OrderProductInfo `json:"orderProductInfo,omitempty"`
	Creator          *pgModel.User       `json:"creator,omitempty"`
	Implementer      *pgModel.User       `json:"implementer,omitempty"`
}

// OrderProductInfo struct
type OrderProductInfo struct {
	OrderQuantity int              `json:"orderQuantity,omitempty"`
	RealPrice     int              `json:"realPrice,omitempty"`
	Product       *pgModel.Product `json:"product,omitempty"`
}

// GetOrderResponse struct
type GetOrderResponse struct {
	Order *Order `json:"order,omitempty"`
}

// GetOrdersResponse struct
type GetOrdersResponse struct {
	Orders []*Order `json:"orders"`
}

// UpdateOrderResponse struct
type UpdateOrderResponse struct {
	Order         *Order  `json:"order,omitempty"`
	TransactionID *string `json:"transactionID,omitempty"`
}
