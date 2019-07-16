package user

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	domainModel "github.com/phungvandat/life-cafe-backend/model/domain"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/config"
	"github.com/phungvandat/life-cafe-backend/util/helper"
)

// pgService implmenter for User serivce in postgres
type pgService struct {
	db     *gorm.DB
	logger log.Logger
}

// NewPGService new pg service
func NewPGService(db *gorm.DB, logger log.Logger) Service {
	return &pgService{
		db:     db,
		logger: logger,
	}
}

func (s *pgService) Create(ctx context.Context, req requestModel.CreateUserRequest) (*responseModel.CreateUserResponse, error) {
	userExisted := &domainModel.User{Username: req.User.Username}
	err := s.db.Find(userExisted, userExisted).Error
	if err == nil {
		return nil, UsernameIsExistedError
	}

	user := &domainModel.User{
		Username: req.User.Username,
		Fullname: req.User.Fullname,
		Password: req.User.Password,
		Role:     req.User.Role,
		Active:   req.User.Active,
	}
	err = s.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return &responseModel.CreateUserResponse{
		User: user,
	}, nil
}

func (s *pgService) LogIn(ctx context.Context, req requestModel.UserLogInRequest) (*responseModel.UserLogInResponse, error) {
	username := req.Username
	password := req.Password

	user := &domainModel.User{Username: username}
	err := s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, UserNotFoundError
	}
	checkPassword := user.ComparePassword(password)

	if checkPassword == false {
		return nil, WrongPasswordError
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
	user := &domainModel.User{
		Username: "master",
	}
	err := s.db.Find(user, user).Error

	if err == gorm.ErrRecordNotFound {
		user = &domainModel.User{
			Username: "master",
			Password: "master",
			Fullname: "master",
			Role:     "admin",
		}
		return s.db.Create(user).Error
	}
	return err
}
