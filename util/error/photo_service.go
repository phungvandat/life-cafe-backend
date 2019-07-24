package error

import (
	"net/http"
)

//Error Declaration
var (
	InvalidPhotoURLError = invalidPhotoURLError{}
	InvalidPhotoIDError  = invalidPhotoIDError{}
)

type invalidPhotoURLError struct{}

func (invalidPhotoURLError) Error() string {
	return "Invalid photo URL"
}

func (invalidPhotoURLError) StatusCode() int {
	return http.StatusBadRequest
}

type invalidPhotoIDError struct{}

func (invalidPhotoIDError) Error() string {
	return "Invalid photo ID"
}

func (invalidPhotoIDError) StatusCode() int {
	return http.StatusBadRequest
}
