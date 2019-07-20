package auth

import ("context")

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

func (mw validationMiddleware) AuthenticateUser(ctx context.Context) error {
	return mw.Service.AuthenticateUser(ctx)
}

func (mw validationMiddleware) AuthenticateAdmin(ctx context.Context) error {
	return mw.Service.AuthenticateAdmin(ctx)
}
