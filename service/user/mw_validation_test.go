package user

import (
	"context"
	"reflect"
	"testing"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
)

func Test_validationMiddleware_Create(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	serviceMock := &ServiceMock{
		CreateFunc: func(ctx context.Context, _ requestModel.CreateUserRequest) (*responseModel.CreateUserResponse, error) {
			return &responseModel.CreateUserResponse{}, nil
		},
	}

	type args struct {
		ctx context.Context
		req requestModel.CreateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *responseModel.CreateUserResponse
		wantErr error
	}{
		{
			name: "Valid create user",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{
					Username:    "phungvandat",
					Fullname:    "Phung Van Dat",
					Password:    "password",
					PhoneNumber: "0969000111",
					Role:        "master",
				},
			},
			want:    &responseModel.CreateUserResponse{},
			wantErr: nil,
		},
		{
			name: "Failed by missing username",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{},
			},
			want:    nil,
			wantErr: errors.MissingUsernameError,
		},
		{
			name: "Failed by invalid username",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{
					Username: "phung van dat",
				},
			},
			want:    nil,
			wantErr: errors.InvalidUsernameError,
		},
		{
			name: "Failed by missing fullname",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{
					Username: "phungvandat",
				},
			},
			want:    nil,
			wantErr: errors.MissingFullnameError,
		},
		{
			name: "Failed by missing role",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{
					Username: "phungvandat",
					Fullname: "Phung Van Dat",
				},
			},
			want:    nil,
			wantErr: errors.MissingRoleError,
		},
		{
			name: "Failed by invalid role",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{
					Username: "phungvandat",
					Fullname: "Phung Van Dat",
					Role:     "hihi",
				},
			},

			want:    nil,
			wantErr: errors.InvalidRoleError,
		},
		{
			name: "Failed by missing password",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{
					Username: "phungvandat",
					Fullname: "Phung Van Dat",
					Role:     "master",
				},
			},

			want:    nil,
			wantErr: errors.MissingPasswordError,
		},
		{
			name: "Failed by missing phone number",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{
					Username: "phungvandat",
					Fullname: "Phung Van Dat",
					Role:     "master",
					Password: "huhuhuhu",
				},
			},

			want:    nil,
			wantErr: errors.UserPhoneNumberIsRequiredError,
		},
		{
			name: "Failed by invalid phone number",
			args: args{
				ctx: ctx,
				req: requestModel.CreateUserRequest{
					Username:    "phungvandat",
					Fullname:    "Phung Van Dat",
					Role:        "master",
					Password:    "huhuhuhu",
					PhoneNumber: "12345a",
				},
			},

			want:    nil,
			wantErr: errors.InvalidUserPhoneNumberError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			got, err := mw.Create(tt.args.ctx, tt.args.req)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
