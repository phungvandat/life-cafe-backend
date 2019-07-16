package middlewares

import (
	"github.com/phungvandat/life-cafe-backend/http/middlewares/auth"
	"github.com/phungvandat/life-cafe-backend/service"
)

// Middlewares struct
type Middlewares struct {
	AuthMiddleware auth.AuthMiddleware
}

// MakeHTTPpMiddleware func
func MakeHTTPpMiddleware(s service.Service) Middlewares {
	return Middlewares{
		AuthMiddleware: auth.AuthMiddleware{
			AuthUser:  auth.MakeAuthUserMiddleware(s),
			AuthAdmin: auth.MakeAuthAdminMiddleware(s),
		},
	}
}
