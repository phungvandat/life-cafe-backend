package auth

import (
	"context"
)

// Service interface contain functions
type Service interface {
	AuthenticateUser(ctx context.Context) error
	AuthenticateAdmin(ctx context.Context) error
}
