package request

// CreateOrderRequest struct
type CreateOrderRequest struct {
	// Order information
	Type                string `json:"type,omitempty"`
	Note                string `json:"note,omitempty"`
	CustomerID          string `json:"customerID,omitempty"`
	Status              string `json:"status,omitempty"`
	ReceiverPhoneNumber string `json:"receiverPhoneNumber,omitempty"`
	ReceiverAddress     string `json:"receiverAddress,omitempty"`
	ReceiverFullname    string `json:"receiverFullname,omitempty"`
	// Product information
	OrderProductInfo []ProductOrder `json:"orderProductInfo,omitempty"`
	// Customer information
	CustomerPhoneNumber string `json:"customerPhoneNumber,omitempty"`
	CustomerAddress     string `json:"customerAddress,omitempty"`
	CustomerFullname    string `json:"customerFullname,omitempty"`
}

// ProductOrder struct
type ProductOrder struct {
	ProductID      string `json:"productID,omitempty"`
	OrderQuantity  int    `json:"orderQuantity,omitempty"`
	OrderRealPrice int    `json:"orderRealPrice,omitempty"`
}

// GetOrderRequest struct
type GetOrderRequest struct {
	ParamOrderID string `json:"orderID,omitempty"`
}

// GetOrdersRequest struct
type GetOrdersRequest struct {
	Skip  string `json:"skip,omitempty"`
	Limit string `json:"limit,omitempty"`
}

// UpdateOrderRequest struct
type UpdateOrderRequest struct {
	ParamOrderID string `json:"orderID,omitempty"`
	// Order information
	Note                string `json:"note,omitempty"`
	Status              string `json:"status,omitempty"`
	ReceiverPhoneNumber string `json:"receiverPhoneNumber,omitempty"`
	ReceiverAddress     string `json:"receiverAddress,omitempty"`
	ReceiverFullname    string `json:"receiverFullname,omitempty"`
	// Product information
	OrderProductInfo []ProductOrder `json:"orderProductInfo,omitempty"`
}
