package response

import (
	domainModel "github.com/phungvandat/life-cafe-backend/model/domain"
)

// CreateUserResponse struct
type CreateUserResponse struct {
	User *domainModel.User `json:"user,omitempty"`
}

// UserLogInResponse struct
type UserLogInResponse struct {
	User  *domainModel.User `json:"user,omitempty"`
	Token string            `json:"token,omitempty"`
}
