package user

import (
	"github.com/go-kit/kit/endpoint"
)

// UserEndpoint struct
type UserEndpoint struct {
	CreateUser endpoint.Endpoint
	UserLogin  endpoint.Endpoint
	GetUsers   endpoint.Endpoint
	GetUser    endpoint.Endpoint
}
