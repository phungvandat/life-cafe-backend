package error

import (
	"net/http"
)

//Error Declaration
var (
	MissingUsernameError           = missingUsernameError{}
	MissingFullnameError           = missingFullnameError{}
	MissingPasswordError           = missingPasswordError{}
	MissingRoleError               = missingRoleError{}
	UserNotFoundError              = userNotFoundError{}
	WrongPasswordError             = wrongPasswordError{}
	UsernameIsExistedError         = usernameIsExistedError{}
	InvalidRoleError               = invalidRoleError{}
	InvalidUsernameError           = invalidUsernameError{}
	UserPhoneNumberIsRequiredError = userPhoneNumberIsRequiredError{}
	InvalidUserPhoneNumberError    = invalidUserPhoneNumberError{}
	PermissionDeniedError          = permissionDeniedError{}
	InvalidUserIDError             = invalidUserIDError{}
)

// Error missing username
type missingUsernameError struct{}

func (missingUsernameError) Error() string {
	return "username is required"
}

func (missingUsernameError) StatusCode() int {
	return http.StatusBadRequest
}

// Error missing fullname
type missingFullnameError struct{}

func (missingFullnameError) Error() string {
	return "fullname is required"
}

func (missingFullnameError) StatusCode() int {
	return http.StatusBadRequest
}

// Error missing password
type missingPasswordError struct{}

func (missingPasswordError) Error() string {
	return "Password is required"
}

func (missingPasswordError) StatusCode() int {
	return http.StatusBadRequest
}

// Error missing role
type missingRoleError struct{}

func (missingRoleError) Error() string {
	return "Role is required"
}

func (missingRoleError) StatusCode() int {
	return http.StatusBadRequest
}

// Error user not found
type userNotFoundError struct{}

func (userNotFoundError) Error() string {
	return "User not found"
}

func (userNotFoundError) StatusCode() int {
	return http.StatusNotFound
}

// Wrong password error
type wrongPasswordError struct{}

func (wrongPasswordError) Error() string {
	return "Wrong password"
}

func (wrongPasswordError) StatusCode() int {
	return http.StatusUnauthorized
}

// Username existed error
type usernameIsExistedError struct{}

func (usernameIsExistedError) Error() string {
	return "Username is existed"
}

func (usernameIsExistedError) StatusCode() int {
	return http.StatusBadRequest
}

// invalid role error
type invalidRoleError struct{}

func (invalidRoleError) Error() string {
	return "Invalid role user"
}

func (invalidRoleError) StatusCode() int {
	return http.StatusBadRequest
}

// Invalid username error
type invalidUsernameError struct{}

func (invalidUsernameError) Error() string {
	return "Invalid username"
}

func (invalidUsernameError) StatusCode() int {
	return http.StatusBadRequest
}

//
type userPhoneNumberIsRequiredError struct{}

func (userPhoneNumberIsRequiredError) Error() string {
	return "User phone number is required"
}

func (userPhoneNumberIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

//
type invalidUserPhoneNumberError struct{}

func (invalidUserPhoneNumberError) Error() string {
	return "Invalid user phone number"
}

func (invalidUserPhoneNumberError) StatusCode() int {
	return http.StatusBadRequest
}

//
type permissionDeniedError struct{}

func (permissionDeniedError) Error() string {
	return "Permission denied"
}

func (permissionDeniedError) StatusCode() int {
	return http.StatusForbidden
}

//
type invalidUserIDError struct{}

func (invalidUserIDError) Error() string {
	return "Invalid user ID"
}

func (invalidUserIDError) StatusCode() int {
	return http.StatusBadRequest
}
