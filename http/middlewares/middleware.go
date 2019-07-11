package middlewares

import (
	"github.com/phungvandat/life-cafe-backend/http/middlewares/auth"
	"github.com/phungvandat/life-cafe-backend/service"
)

type Middlewares struct {
	auth.AuthMiddleware
}

func MakeHttpMiddleware(s service.Service) Middlewares {
	return Middlewares{
		AuthMiddleware: auth.AuthMiddleware{
			AuthUser: auth.MakeAuthUserMiddleware(s),
		},
	}
}
