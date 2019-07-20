package error

import (
	"net/http"
)

//Error Declaration
var (
	ProductNameIsRequiredError         = productNameIsRequiredError{}
	ProductMainPhotoIsRequiredError    = productMainPhotoIsRequiredError{}
	InvalidProductQuantityError        = invalidProductQuantityError{}
	CategoryOfProductIsRequiredError   = categoryOfProductIsRequiredError{}
	DuplicateCategoryError             = duplicateCategoryError{}
	InvalidProductPriceError           = invalidProductPriceError{}
	ProductSlugIsRequiredError         = productSlugIsRequiredError{}
	InvalidProductSlugError            = invalidProductSlugError{}
	InvalidCategoryIDError             = invalidCategoryIDError{}
	ProductSlugExistError              = productSlugExistError{}
	InvalidSecondaryPhotoQuantityError = invalidSecondaryPhotoQuantityError{}
	InvalidProductIDTypeError          = invalidProductIDTypeError{}
	ProductNotExistError               = productNotExistError{}
	InvalidProductNameError            = invalidProductNameError{}
	ProductIDIsRequiredError           = productIDIsRequiredError{}
)

//
type productNameIsRequiredError struct{}

func (productNameIsRequiredError) Error() string {
	return "Product name is requied"
}

func (productNameIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type productMainPhotoIsRequiredError struct{}

func (productMainPhotoIsRequiredError) Error() string {
	return "Product main photo is required"
}

func (productMainPhotoIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidProductQuantityError struct{}

func (invalidProductQuantityError) Error() string {
	return "Invalid product quantity"
}

func (invalidProductQuantityError) StatusCode() int {
	return http.StatusBadRequest
}

//
type categoryOfProductIsRequiredError struct{}

func (categoryOfProductIsRequiredError) Error() string {
	return "Category of product required"
}

func (categoryOfProductIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type duplicateCategoryError struct{}

func (duplicateCategoryError) Error() string {
	return "Duplicate category"
}

func (duplicateCategoryError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidProductPriceError struct{}

func (invalidProductPriceError) Error() string {
	return "Invalid price product"
}

func (invalidProductPriceError) StatusCode() int {
	return http.StatusBadRequest
}

//
type productSlugIsRequiredError struct{}

func (productSlugIsRequiredError) Error() string {
	return "Product slug is required"
}

func (productSlugIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidProductSlugError struct{}

func (invalidProductSlugError) Error() string {
	return "Invalid product slug"
}

func (invalidProductSlugError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidCategoryIDError struct{}

func (invalidCategoryIDError) Error() string {
	return "Invalid category ID"
}

func (invalidCategoryIDError) StatusCode() int {
	return http.StatusBadRequest
}

//
type productSlugExistError struct{}

func (productSlugExistError) Error() string {
	return "Product slug exist"
}

func (productSlugExistError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidSecondaryPhotoQuantityError struct{}

func (invalidSecondaryPhotoQuantityError) Error() string {
	return "Invalid secondary photo quantity"
}

func (invalidSecondaryPhotoQuantityError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidProductIDTypeError struct{}

func (invalidProductIDTypeError) Error() string {
	return "Invalid product ID"
}

func (invalidProductIDTypeError) StatusCode() int {
	return http.StatusBadRequest
}

//
type productNotExistError struct{}

func (productNotExistError) Error() string {
	return "Product not exist"
}

func (productNotExistError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidProductNameError struct{}

func (invalidProductNameError) Error() string {
	return "Invalid product name"
}

func (invalidProductNameError) StatusCode() int {
	return http.StatusBadRequest
}

//
type productIDIsRequiredError struct{}

func (productIDIsRequiredError) Error() string {
	return "Product ID is required"
}

func (productIDIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}
