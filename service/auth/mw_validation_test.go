package auth

import (
	"context"
	"testing"
)

func Test_validationMiddleware_AuthenticateUser(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			if err := mw.AuthenticateUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validationMiddleware_AuthenticateAdmin(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			if err := mw.AuthenticateAdmin(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.AuthenticateAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
