package error

import (
	"net/http"
)

//Error Declaration
var (
	ParentCategoryIsNotExistError    = parentCategoryIsNotExistError{}
	CategorySlugAlreadyExistError    = categorySlugAlreadyExistError{}
	CategoryNameIsRequiredError      = categoryNameIsRequiredError{}
	InvalidParentCategoryIDTypeError = invalidParentCategoryIDTypeError{}
	CategorySlugIsRequiredError      = categorySlugIsRequiredError{}
	InvalidCategorySlugError         = invalidCategorySlugError{}
	CategoryNotExistError            = categoryNotExistError{}
	InvalidCategoryIDTypeError       = invalidCategoryIDTypeError{}
	InvalidSkipError                 = invalidSkipError{}
	InvalidLimitError                = invalidLimitError{}
	InvalidCategoryNameError         = invalidCategoryNameError{}
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
type categorySlugAlreadyExistError struct{}

func (categorySlugAlreadyExistError) Error() string {
	return "Slug already exist"
}

func (categorySlugAlreadyExistError) StatusCode() int {
	return http.StatusBadRequest
}

// Error name required
type categoryNameIsRequiredError struct{}

func (categoryNameIsRequiredError) Error() string {
	return "Category name is required"
}

func (categoryNameIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

// Error invalid parent category id
type invalidParentCategoryIDTypeError struct{}

func (invalidParentCategoryIDTypeError) Error() string {
	return "Invalid parent category ID"
}

func (invalidParentCategoryIDTypeError) StatusCode() int {
	return http.StatusBadRequest
}

//
type categorySlugIsRequiredError struct{}

func (categorySlugIsRequiredError) Error() string {
	return "Slug is required"
}

func (categorySlugIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidCategorySlugError struct{}

func (invalidCategorySlugError) Error() string {
	return "Invalid category slug"
}

func (invalidCategorySlugError) StatusCode() int {
	return http.StatusBadRequest
}

//
type categoryNotExistError struct{}

func (categoryNotExistError) Error() string {
	return "Category not exist"
}

func (categoryNotExistError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidCategoryIDTypeError struct{}

func (invalidCategoryIDTypeError) Error() string {
	return "Invalid category ID"
}

func (invalidCategoryIDTypeError) StatusCode() int {
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
type invalidCategoryNameError struct{}

func (invalidCategoryNameError) Error() string {
	return "Invalid category name"
}

func (invalidCategoryNameError) StatusCode() int {
	return http.StatusBadRequest
}
