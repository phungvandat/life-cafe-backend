package service

import (
	"github.com/phungvandat/life-cafe-backend/service/auth"
	"github.com/phungvandat/life-cafe-backend/service/user"
)

// Service define list of all services in project
type Service struct {
	UserService user.Service
	AuthService auth.Service
}
