package user

import (
	"github.com/phungvandat/life-cafe-backend/domain"
)

// CreateRequest struct
type CreateRequest struct {
	domain.User `json:"user,omitempty"`
}

// CreateResponse struct
type CreateResponse struct {
	*domain.User `json:"user,omitempty"`
}

// LogInRequest struct
type LogInRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// LogInResponse struct
type LogInResponse struct {
	*domain.User `json:"user.omitempty"`
	Token        string `json:"token,omitempty"`
}
