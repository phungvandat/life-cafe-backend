package response

import (
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
)

// CreateUserResponse struct
type CreateUserResponse struct {
	User          *pgModel.User `json:"user,omitempty"`
	TransactionID *string       `json:"transactionID,omitempty"`
}

// UserLogInResponse struct
type UserLogInResponse struct {
	User  *pgModel.User `json:"user,omitempty"`
	Token string        `json:"token,omitempty"`
}

// GetUserResponse struct
type GetUserResponse struct {
	User *pgModel.User `json:"user,omitempty"`
}

// GetUsersResponse struct
type GetUsersResponse struct {
	Users []*pgModel.User `json:"users"`
}
