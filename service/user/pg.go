package user

import (
	"context"

	"github.com/jinzhu/gorm"
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/config"
	"github.com/phungvandat/life-cafe-backend/util/contextkey"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
	"github.com/phungvandat/life-cafe-backend/util/helper"
)

// pgService implmenter for User serivce in postgres
type pgService struct {
	db         *gorm.DB
	spRollback helper.SagasService
}

// NewPGService new pg service
func NewPGService(db *gorm.DB, spRollback helper.SagasService) Service {
	return &pgService{
		db:         db,
		spRollback: spRollback,
	}
}

func (s *pgService) RollbackTransaction(_ context.Context, transactionID string) error {
	return s.spRollback.RollbackTransaction(transactionID)
}

func (s *pgService) CommitTransaction(_ context.Context, transactionID string) error {
	return s.spRollback.CommitTransaction(transactionID)
}

func (s *pgService) Create(ctx context.Context, req requestModel.CreateUserRequest) (*responseModel.CreateUserResponse, error) {
	tx := s.db.Begin()
	transactionID := (pgModel.NewUUID()).String()
	s.spRollback.NewTransaction(transactionID, tx)
	res := &responseModel.CreateUserResponse{
		TransactionID: &transactionID,
	}

	userExisted := &pgModel.User{Username: req.Username}
	err := tx.Find(userExisted, userExisted).Error
	if err == nil {
		return res, errors.UsernameIsExistedError
	}

	user := &pgModel.User{
		Username:    req.Username,
		Fullname:    req.Fullname,
		Password:    req.Password,
		Role:        req.Role,
		Active:      req.Active,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Email:       req.Email,
	}
	err = tx.Create(user).Error
	if err != nil {
		return res, err
	}
	res.User = user

	return res, nil
}

func (s *pgService) LogIn(ctx context.Context, req requestModel.UserLogInRequest) (*responseModel.UserLogInResponse, error) {
	username := req.Username
	password := req.Password

	user := &pgModel.User{Username: username}
	err := s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, errors.UserNotFoundError
	}
	checkPassword := user.ComparePassword(password)

	if checkPassword == false {
		return nil, errors.WrongPasswordError
	}
	clasms := helper.TokenClaims{
		UserID:   user.Model.ID.String(),
		Username: user.Username,
		Role:     user.Role,
	}

	jwt, err := helper.GenerateToken(config.GetJWTSerectKeyEnv(), clasms)

	if err != nil {
		return nil, err
	}

	return &responseModel.UserLogInResponse{
		User:  user,
		Token: jwt,
	}, err
}

func (s *pgService) CreateMaster(_ context.Context) error {
	user := &pgModel.User{
		Username: "master",
	}
	err := s.db.Find(user, user).Error

	if err == gorm.ErrRecordNotFound {
		user = &pgModel.User{
			Username: "master",
			Password: "master",
			Fullname: "master",
			Role:     "admin",
		}
		return s.db.Create(user).Error
	}
	return err
}

func (s *pgService) GetUser(ctx context.Context, req requestModel.GetUserRequest) (*responseModel.GetUserResponse, error) {
	res := &responseModel.GetUserResponse{}

	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)

	userSiginRole := ctx.Value(contextkey.UserRoleContextKey).(string)

	if !check {
		return res, errors.NotLoggedInError
	}

	if userSiginRole != "admin" && userSiginRole != "master" && ctxUserID != req.ParamUserID {
		return res, errors.PermissionDeniedError
	}

	userIDUUID, _ := pgModel.UUIDFromString(req.ParamUserID)
	user := &pgModel.User{
		Model: pgModel.Model{
			ID: userIDUUID,
		},
	}

	err := s.db.Find(user, user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.UserNotFoundError
		}
		return res, err
	}

	res.User = user

	return res, nil
}

func (s *pgService) GetUsers(ctx context.Context, req requestModel.GetUsersRequest) (*responseModel.GetUsersResponse, error) {
	res := &responseModel.GetUsersResponse{}

	skip := req.Skip
	limit := req.Limit
	if req.Skip == "" {
		skip = "-1"
	}

	if req.Limit == "" {
		limit = "-1"
	}

	stringQuery := ""
	if req.Fullname != "" {
		query := "fullname = '" + req.Fullname + "'"
		if stringQuery == "" {
			stringQuery += query
		} else {
			stringQuery += " AND " + query
		}
	}

	if req.AlwaysPhone == "true" {
		query := "phone_number = '" + req.PhoneNumber + "'"
		if stringQuery == "" {
			stringQuery += query
		} else {
			stringQuery += " AND " + query
		}
	} else if req.PhoneNumber != "" {
		query := "phone_number = '" + req.PhoneNumber + "'"
		if stringQuery == "" {
			stringQuery += query
		} else {
			stringQuery += " AND " + query
		}
	}

	users := []*pgModel.User{}

	err := s.db.Limit(limit).Offset(skip).Where(stringQuery).Table("users").Scan(&users).Error

	if err != nil {
		return res, err
	}

	res.Users = users

	return res, nil
}
