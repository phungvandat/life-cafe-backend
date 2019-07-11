package user

import (
	"github.com/go-kit/kit/endpoint"
)

type UserEndpoint struct {
	CreateUser endpoint.Endpoint
	UserLogin  endpoint.Endpoint
}
