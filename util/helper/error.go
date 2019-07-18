package helper

import (
	"net/http"
)

//Error Declaration
var (
	InvalidSigningAlgorithm           = invalidSigningAlgorithm{}
	InvalidAccessToken                = invalidAccessToken{}
	TransactionRollbackeNotExistError = transactionRollbackeNotExistError{}
)

// Error signing algorithm
type invalidSigningAlgorithm struct{}

func (invalidSigningAlgorithm) Error() string {
	return "Invalid signing algorithm"
}

func (invalidSigningAlgorithm) StatusCode() int {
	return http.StatusInternalServerError
}

// Error invalid access token
type invalidAccessToken struct{}

func (invalidAccessToken) Error() string {
	return "Invalid access token"
}

func (invalidAccessToken) StatusCode() int {
	return http.StatusBadRequest
}

//
type transactionRollbackeNotExistError struct{}

func (transactionRollbackeNotExistError) Error() string {
	return "Transaction rollback not exist"
}

func (transactionRollbackeNotExistError) StatusCode() int {
	return http.StatusInternalServerError
}
