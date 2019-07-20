package error

import (
	"net/http"
)

//Error Declaration
var (
	OrderTypeIsRequiredError           = orderTypeIsRequiredError{}
	InvalidOrderTypeError              = invalidOrderTypeError{}
	OrderStatusIsRequiredError         = orderStatusIsRequiredError{}
	InvalidOrderStatusError            = invalidOrderStatusError{}
	OrderProductInfoIsRequiredError    = orderProductInfoIsRequiredError{}
	InvalidOrderProductQuantityError   = invalidOrderProductQuantityError{}
	InvalidCustomerIDError             = invalidCustomerIDError{}
	ReceiverPhoneNumberIsRequiredError = receiverPhoneNumberIsRequiredError{}
	ReceiverAddressIsRequiredError     = receiverAddressIsRequiredError{}
	ReceiverFullnameIsRequiredError    = receiverFullnameIsRequiredError{}
	InvalidReceiverPhoneNumberError    = invalidReceiverPhoneNumberError{}
	InvalidOrderProductRealPriceError  = invalidOrderProductRealPriceError{}
	CustomerPhoneNumberIsReqiredError  = customerPhoneNumberIsReqiredError{}
	InvalidCustomerPhoneNumberError    = invalidCustomerPhoneNumberError{}
	CustomerAddressIsRequiredError     = customerAddressIsRequiredError{}
	CustomerFullnameIsRequiredError    = customerFullnameIsRequiredError{}
	CustomerNotExistError              = customerNotExistError{}
	OrderNotExistError                 = orderNotExistError{}
	CannotUpdateOrderError             = cannotUpdateOrderError{}
)

//
type invalidOrderTypeError struct{}

func (invalidOrderTypeError) Error() string {
	return "Invalid order type"
}

func (invalidOrderTypeError) StatusCode() int {
	return http.StatusBadRequest
}

//
type orderTypeIsRequiredError struct{}

func (orderTypeIsRequiredError) Error() string {
	return "Order type is required"
}

func (orderTypeIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type orderStatusIsRequiredError struct{}

func (orderStatusIsRequiredError) Error() string {
	return "Order status is required"
}

func (orderStatusIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidOrderStatusError struct{}

func (invalidOrderStatusError) Error() string {
	return "Invalid order status"
}

func (invalidOrderStatusError) StatusCode() int {
	return http.StatusBadRequest
}

//
type orderProductInfoIsRequiredError struct{}

func (orderProductInfoIsRequiredError) Error() string {
	return "Order product info is required"
}

func (orderProductInfoIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidOrderProductQuantityError struct{}

func (invalidOrderProductQuantityError) Error() string {
	return "Invalid order product quantity"
}

func (invalidOrderProductQuantityError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidCustomerIDError struct{}

func (invalidCustomerIDError) Error() string {
	return "Invalid customer ID"
}

func (invalidCustomerIDError) StatusCode() int {
	return http.StatusBadRequest
}

//
type receiverPhoneNumberIsRequiredError struct{}

func (receiverPhoneNumberIsRequiredError) Error() string {
	return "Receiver phone number is required"
}

func (receiverPhoneNumberIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type receiverAddressIsRequiredError struct{}

func (receiverAddressIsRequiredError) Error() string {
	return "Receiver address is required"
}

func (receiverAddressIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type receiverFullnameIsRequiredError struct{}

func (receiverFullnameIsRequiredError) Error() string {
	return "Receiver fullname is required"
}

func (receiverFullnameIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidReceiverPhoneNumberError struct{}

func (invalidReceiverPhoneNumberError) Error() string {
	return "Invalid receiver phone number"
}

func (invalidReceiverPhoneNumberError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidOrderProductRealPriceError struct{}

func (invalidOrderProductRealPriceError) Error() string {
	return "Invalid order product real price"
}

func (invalidOrderProductRealPriceError) StatusCode() int {
	return http.StatusBadRequest
}

//
type customerPhoneNumberIsReqiredError struct{}

func (customerPhoneNumberIsReqiredError) Error() string {
	return "Customer phone number is required"
}

func (customerPhoneNumberIsReqiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidCustomerPhoneNumberError struct{}

func (invalidCustomerPhoneNumberError) Error() string {
	return "Invalid customer phone number"
}

func (invalidCustomerPhoneNumberError) StatusCode() int {
	return http.StatusBadRequest
}

//
type customerAddressIsRequiredError struct{}

func (customerAddressIsRequiredError) Error() string {
	return "Customer address is required"
}

func (customerAddressIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type customerFullnameIsRequiredError struct{}

func (customerFullnameIsRequiredError) Error() string {
	return "Customer fullname is required"
}

func (customerFullnameIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type customerNotExistError struct{}

func (customerNotExistError) Error() string {
	return "Customer not exist"
}

func (customerNotExistError) StatusCode() int {
	return http.StatusBadRequest
}

//
type orderNotExistError struct{}

func (orderNotExistError) Error() string {
	return "Order not exist"
}

func (orderNotExistError) StatusCode() int {
	return http.StatusBadRequest
}

//
type cannotUpdateOrderError struct{}

func (cannotUpdateOrderError) Error() string {
	return "Cannot update order"
}

func (cannotUpdateOrderError) StatusCode() int {
	return http.StatusBadRequest
}
