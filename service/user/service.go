package user

import (
	"context"

	"github.com/phungvandat/life-cafe-backend/domain"
)

// Service interface contain functions
type Service interface {
	Create(ctx context.Context, u domain.User) (*domain.User, error)
	LogIn(ctx context.Context, u domain.User) (*domain.User, string, error)
}
