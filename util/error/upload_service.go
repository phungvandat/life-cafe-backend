package error

import (
	"net/http"
)

//Error Declaration
var (
	CreateFileError      = createFileError{}
	CopyFileError        = copyFileError{}
	FileAreNotImageTypes = filesAreNotImageTypes{}
)

// Create file error
type createFileError struct{}

func (createFileError) Error() string {
	return "Create file failure"
}

func (createFileError) StatusCode() int {
	return http.StatusInternalServerError
}

// Copy file error
type copyFileError struct{}

func (copyFileError) Error() string {
	return "Copy file error"
}

func (copyFileError) StatusCode() int {
	return http.StatusInternalServerError
}

// Files not image type error

type filesAreNotImageTypes struct{}

func (filesAreNotImageTypes) Error() string {
	return "Files are not image types"
}

func (filesAreNotImageTypes) StatusCode() int {
	return http.StatusBadRequest
}
