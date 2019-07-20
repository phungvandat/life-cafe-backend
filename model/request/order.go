package request

// CreateOrderRequest struct
type CreateOrderRequest struct {
	// Order information
	Type                string `json:"type,omitempty"`
	Note                string `json:"note,omitempty"`
	CustomerID          string `json:"customer_id,omitempty"`
	Status              string `json:"status,omitempty"`
	ReceiverPhoneNumber string `json:"receiver_phone_number,omitempty"`
	ReceiverAddress     string `json:"receiver_address,omitempty"`
	ReceiverFullname    string `json:"receiver_fullname,omitempty"`
	// Product information
	OrderProductInfo []ProductOrder `json:"order_product_info,omitempty"`
	// Customer information
	CustomerPhoneNumber string `json:"customer_phone_number,omitempty"`
	CustomerAddress     string `json:"customer_address,omitempty"`
	CustomerFullname    string `json:"customer_fullname,omitempty"`
}

// ProductOrder struct
type ProductOrder struct {
	ProductID     string `json:"product_id,omitempty"`
	OrderQuantity int    `json:"order_quantity,omitempty"`
	RealPrice     int    `json:"real_price,omitempty"`
}

// GetOrderRequest struct
type GetOrderRequest struct {
	ParamOrderID string `json:"order_id,omitempty"`
}

// GetOrdersRequest struct
type GetOrdersRequest struct {
	Skip  string `json:"skip,omitempty"`
	Limit string `json:"limit,omitempty"`
}

// UpdateOrderRequest struct
type UpdateOrderRequest struct {
	ParamOrderID string `json:"order_id,omitempty"`
	// Order information
	Note                string `json:"note,omitempty"`
	Status              string `json:"status,omitempty"`
	ReceiverPhoneNumber string `json:"receiver_phone_number,omitempty"`
	ReceiverAddress     string `json:"receiver_address,omitempty"`
	ReceiverFullname    string `json:"receiver_fullname,omitempty"`
	// Product information
	OrderProductInfo []ProductOrder `json:"order_product_info,omitempty"`
}
