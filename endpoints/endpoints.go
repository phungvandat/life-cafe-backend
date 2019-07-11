package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phungvandat/life-cafe-backend/endpoints/index"
	"github.com/phungvandat/life-cafe-backend/endpoints/user"
	"github.com/phungvandat/life-cafe-backend/service"
)

type Endpoints struct {
	Index endpoint.Endpoint
	user.UserEndpoint
}

func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Index: index.MakeIndexEndpoints(),
		UserEndpoint: user.UserEndpoint{
			CreateUser: user.MakeCreateEndpoint(s),
			UserLogin:  user.MakeLogInEndpoint(s),
		},
	}
}
