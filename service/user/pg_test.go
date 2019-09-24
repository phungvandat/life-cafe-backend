package user

import (
	"context"
	"reflect"
	"testing"

	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/config/db/pg"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
	"github.com/phungvandat/life-cafe-backend/util/externals/sagas"
)

func Test_pgService_Create(t *testing.T) {
	t.Parallel()
	testDB, cleanup := pg.CreateTestDatabase(t)
	defer cleanup()

	err := pg.MigrateTables(testDB)

	if err != nil {
		t.Fatalf("Failed to migrate test table by error %v", err)
	}

	// Create user to failed by username exist
	existedUser := pgModel.User{
		Username: "phungvandat",
		Fullname: "Phung Van Dat",
	}

	err = testDB.Create(&existedUser).Error
	if err != nil {
		t.Fatalf("Failed create existed user by error %v", err)
	}

	type args struct {
		req requestModel.CreateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *responseModel.CreateUserResponse
		wantErr error
	}{
		{
			name: "Create user success",
			args: args{
				req: requestModel.CreateUserRequest{
					Username:    "phungvandat2",
					Fullname:    "Phung Van Dat",
					Role:        "master",
					Password:    "Hello",
					PhoneNumber: "0909060606",
					Active:      true,
					Address:     "address",
					Email:       "email",
				},
			},
			want: &responseModel.CreateUserResponse{
				User: &pgModel.User{
					Username:    "phungvandat2",
					Fullname:    "Phung Van Dat",
					Role:        "master",
					PhoneNumber: "0909060606",
					Active:      true,
					Address:     "address",
					Email:       "email",
				},
			},
			wantErr: nil,
		},
		{
			name: "Failed by username existed",
			args: args{
				req: requestModel.CreateUserRequest{
					Username:    "phungvandat",
					Fullname:    "Phung Van Dat",
					Role:        "master",
					Password:    "Hello",
					PhoneNumber: "0909060606",
					Active:      true,
					Address:     "address",
					Email:       "email",
				},
			},
			want:    nil,
			wantErr: errors.UsernameIsExistedError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db:         testDB,
				spRollback: sagas.NewSagasService(),
			}

			got, err := s.Create(context.Background(), tt.args.req)

			if (err != nil) && err != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr %v", err, tt.wantErr)
				s.spRollback.RollbackAllTransaction()
				return
			}

			if tt.want != nil {
				wantUser := &responseModel.CreateUserResponse{
					User: &pgModel.User{
						Username:    got.User.Username,
						Fullname:    got.User.Fullname,
						Role:        got.User.Role,
						PhoneNumber: got.User.PhoneNumber,
						Active:      got.User.Active,
						Address:     got.User.Address,
						Email:       got.User.Email,
					},
				}
				if !reflect.DeepEqual(wantUser, tt.want) {
					t.Errorf("pgService.Create() = %v, want %v", wantUser, tt.want)
				}
			}
			s.spRollback.CommitAllTransaction()
		})
	}
}
