// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package auth

import (
	"context"
	"sync"
)

var (
	lockServiceMockAuthenticateAdmin sync.RWMutex
	lockServiceMockAuthenticateUser  sync.RWMutex
)

// Ensure, that ServiceMock does implement Service.
// If this is not the case, regenerate this file with moq.
var _ Service = &ServiceMock{}

// ServiceMock is a mock implementation of Service.
//
//     func TestSomethingThatUsesService(t *testing.T) {
//
//         // make and configure a mocked Service
//         mockedService := &ServiceMock{
//             AuthenticateAdminFunc: func(ctx context.Context) error {
// 	               panic("mock out the AuthenticateAdmin method")
//             },
//             AuthenticateUserFunc: func(ctx context.Context) error {
// 	               panic("mock out the AuthenticateUser method")
//             },
//         }
//
//         // use mockedService in code that requires Service
//         // and then make assertions.
//
//     }
type ServiceMock struct {
	// AuthenticateAdminFunc mocks the AuthenticateAdmin method.
	AuthenticateAdminFunc func(ctx context.Context) error

	// AuthenticateUserFunc mocks the AuthenticateUser method.
	AuthenticateUserFunc func(ctx context.Context) error

	// calls tracks calls to the methods.
	calls struct {
		// AuthenticateAdmin holds details about calls to the AuthenticateAdmin method.
		AuthenticateAdmin []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// AuthenticateUser holds details about calls to the AuthenticateUser method.
		AuthenticateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
}

// AuthenticateAdmin calls AuthenticateAdminFunc.
func (mock *ServiceMock) AuthenticateAdmin(ctx context.Context) error {
	if mock.AuthenticateAdminFunc == nil {
		panic("ServiceMock.AuthenticateAdminFunc: method is nil but Service.AuthenticateAdmin was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockServiceMockAuthenticateAdmin.Lock()
	mock.calls.AuthenticateAdmin = append(mock.calls.AuthenticateAdmin, callInfo)
	lockServiceMockAuthenticateAdmin.Unlock()
	return mock.AuthenticateAdminFunc(ctx)
}

// AuthenticateAdminCalls gets all the calls that were made to AuthenticateAdmin.
// Check the length with:
//     len(mockedService.AuthenticateAdminCalls())
func (mock *ServiceMock) AuthenticateAdminCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockServiceMockAuthenticateAdmin.RLock()
	calls = mock.calls.AuthenticateAdmin
	lockServiceMockAuthenticateAdmin.RUnlock()
	return calls
}

// AuthenticateUser calls AuthenticateUserFunc.
func (mock *ServiceMock) AuthenticateUser(ctx context.Context) error {
	if mock.AuthenticateUserFunc == nil {
		panic("ServiceMock.AuthenticateUserFunc: method is nil but Service.AuthenticateUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockServiceMockAuthenticateUser.Lock()
	mock.calls.AuthenticateUser = append(mock.calls.AuthenticateUser, callInfo)
	lockServiceMockAuthenticateUser.Unlock()
	return mock.AuthenticateUserFunc(ctx)
}

// AuthenticateUserCalls gets all the calls that were made to AuthenticateUser.
// Check the length with:
//     len(mockedService.AuthenticateUserCalls())
func (mock *ServiceMock) AuthenticateUserCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockServiceMockAuthenticateUser.RLock()
	calls = mock.calls.AuthenticateUser
	lockServiceMockAuthenticateUser.RUnlock()
	return calls
}