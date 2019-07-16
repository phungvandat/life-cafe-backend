package productcategory

import (
	"net/http"
)

//Error Declaration
var (
	ParentCategoryIsNotExistError           = parentCategoryIsNotExistError{}
	ProductCategorySlugAlreadyExistError    = productCategorySlugAlreadyExistError{}
	ProductCategoryNameIsRequiredError      = productCategoryNameIsRequiredError{}
	InvalidParentProductCategoryIDTypeError = invalidParentProductCategoryIDTypeError{}
	ProductCategorySlugIsRequiredError      = productCategorySlugIsRequiredError{}
	InvalidProductCategorySlugError         = invalidProductCategorySlugError{}
	ProductCategoryNotExistError            = productCategoryNotExistError{}
	InvalidProductCategoryIDTypeError       = invalidProductCategoryIDTypeError{}
	InvalidSkipError                        = invalidSkipError{}
	InvalidLimitError                       = invalidLimitError{}
	InvalidProductCategoryNameError         = invalidProductCategoryNameError{}
)

// Error parent category not exist
type parentCategoryIsNotExistError struct{}

func (parentCategoryIsNotExistError) Error() string {
	return "Parent category is not exist"
}

func (parentCategoryIsNotExistError) StatusCode() int {
	return http.StatusBadRequest
}

// Error slug exist
type productCategorySlugAlreadyExistError struct{}

func (productCategorySlugAlreadyExistError) Error() string {
	return "Slug already exist"
}

func (productCategorySlugAlreadyExistError) StatusCode() int {
	return http.StatusBadRequest
}

// Error name required
type productCategoryNameIsRequiredError struct{}

func (productCategoryNameIsRequiredError) Error() string {
	return "Product category name is required"
}

func (productCategoryNameIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

// Error invalid parent product category id
type invalidParentProductCategoryIDTypeError struct{}

func (invalidParentProductCategoryIDTypeError) Error() string {
	return "Invalid parent product category ID"
}

func (invalidParentProductCategoryIDTypeError) StatusCode() int {
	return http.StatusBadRequest
}

//
type productCategorySlugIsRequiredError struct{}

func (productCategorySlugIsRequiredError) Error() string {
	return "Slug is required"
}

func (productCategorySlugIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidProductCategorySlugError struct{}

func (invalidProductCategorySlugError) Error() string {
	return "Invalid product category slug"
}

func (invalidProductCategorySlugError) StatusCode() int {
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
type invalidProductCategoryIDTypeError struct{}

func (invalidProductCategoryIDTypeError) Error() string {
	return "Invalid product category ID"
}

func (invalidProductCategoryIDTypeError) StatusCode() int {
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
type invalidProductCategoryNameError struct{}

func (invalidProductCategoryNameError) Error() string {
	return "Invalid product category name"
}

func (invalidProductCategoryNameError) StatusCode() int {
	return http.StatusBadRequest
}
