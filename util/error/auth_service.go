package error

import (
	"net/http"
)

//Error Declaration
var (
	AccountNotFoundError = accountNotFoundError{}
	AccountIsLockedError = accountIsLockedError{}
	NotLoggedInError     = notLoggedInError{}
	AccessDeniedError    = accessDeniedError{}
)

// account not found error
type accountNotFoundError struct{}

func (accountNotFoundError) Error() string {
	return "Account not found"
}

func (accountNotFoundError) StatusCode() int {
	return http.StatusNotFound
}

// account is locked error
type accountIsLockedError struct{}

func (accountIsLockedError) Error() string {
	return "Account is locked"
}

func (accountIsLockedError) StatusCode() int {
	return http.StatusLocked
}

// Not logged in error
type notLoggedInError struct{}

func (notLoggedInError) Error() string {
	return "Please login to continue"
}

func (notLoggedInError) StatusCode() int {
	return http.StatusUnauthorized
}

//
type accessDeniedError struct{}

func (accessDeniedError) Error() string {
	return "Access denied"
}

func (accessDeniedError) StatusCode() int {
	return http.StatusForbidden
}
