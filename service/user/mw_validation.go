package user

import (
	"context"
	"regexp"
	"strings"

	"github.com/phungvandat/life-cafe-backend/domain"
	"github.com/phungvandat/life-cafe-backend/util/constants"
	"github.com/phungvandat/life-cafe-backend/util/regex"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware for validation purposes
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, u domain.User) (*domain.User, error) {
	if u.Username == "" {
		return nil, MissingUsernameError
	}
	usernameRegex, _ := regexp.Compile(regex.UsernameRegex)
	if !usernameRegex.MatchString(u.Username) {
		return nil, InvalidUsernameError
	}

	if strings.Trim(u.Fullname, " ") == "" {
		return nil, MissingFullnameError
	}

	if strings.Trim(u.Password, " ") == "" {
		return nil, MissingPasswordError
	}

	if strings.Trim(u.Role, " ") == "" {
		return nil, MissingRoleError
	}

	if _, check := constants.USER_ROLE[u.Role]; !check {
		return nil, InvalidRoleError
	}

	return mw.Service.Create(ctx, u)
}

func (mw validationMiddleware) LogIn(ctx context.Context, u domain.User) (*domain.User, string, error) {
	if u.Username == "" {
		return nil, "", MissingUsernameError
	}

	if u.Password == "" {
		return nil, "", MissingPasswordError
	}

	return mw.Service.LogIn(ctx, u)
}
