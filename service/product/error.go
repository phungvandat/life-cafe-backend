package product

import (
	"net/http"
)

//Error Declaration
var (
	ProductNameIsRequiredError         = productNameIsRequiredError{}
	ProductMainPhotoIsRequiredError    = productMainPhotoIsRequiredError{}
	InvalidProductQuantityError        = invalidProductQuantityError{}
	CategoryOfProductIsRequiredError   = categoryOfProductIsRequiredError{}
	DuplicateProductCategoryError      = duplicateProductCategoryError{}
	InvalidProductPriceError           = invalidProductPriceError{}
	ProductSlugIsRequiredError         = productSlugIsRequiredError{}
	InvalidProductSlugError            = invalidProductSlugError{}
	InvalidCategoryIDError             = invalidCategoryIDError{}
	ProductSlugExistError              = productSlugExistError{}
	ProductCategoryNotExistError       = productCategoryNotExistError{}
	InvalidSecondaryPhotoQuantityError = invalidSecondaryPhotoQuantityError{}
	InvalidProductIDTypeError          = invalidProductIDTypeError{}
	ProductNotExistError               = productNotExistError{}
	InvalidSkipError                   = invalidSkipError{}
	InvalidLimitError                  = invalidLimitError{}
	InvalidProductNameError            = invalidProductNameError{}
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
type duplicateProductCategoryError struct{}

func (duplicateProductCategoryError) Error() string {
	return "Duplicate product category"
}

func (duplicateProductCategoryError) StatusCode() int {
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
type productCategoryNotExistError struct{}

func (productCategoryNotExistError) Error() string {
	return "Product category not exist"
}

func (productCategoryNotExistError) StatusCode() int {
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
	return "Invalid product category ID"
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
type invalidSkipError struct{}

func (invalidSkipError) Error() string {
	return "Invalid skip"
}

func (invalidSkipError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidLimitError struct{}

func (invalidLimitError) Error() string {
	return "Invalid limit"
}

func (invalidLimitError) StatusCode() int {
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
