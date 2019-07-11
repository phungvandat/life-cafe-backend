package user

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/phungvandat/life-cafe-backend/domain"
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

func (s *pgService) Create(ctx context.Context, u domain.User) (*domain.User, error) {
	userExisted := &domain.User{Username: u.Username}
	err := s.db.Find(userExisted, userExisted).Error
	if err == nil {
		return nil, UsernameIsExistedError
	}

	err = s.db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *pgService) LogIn(ctx context.Context, u domain.User) (*domain.User, string, error) {
	username := u.Username
	password := u.Password

	user := &domain.User{Username: username}
	err := s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, "", UserNotFoundError
	}
	checkPassword := user.ComparePassword(password)

	if checkPassword == false {
		return nil, "", WrongPasswordError
	}
	clasms := helper.TokenClaims{
		UserID:   user.Model.ID.String(),
		Username: user.Username,
		Role:     user.Role,
	}

	jwt, err := helper.GenerateToken(config.GetJWTSerectKeyEnv(), clasms)

	if err != nil {
		return nil, "", err
	}

	return user, jwt, err
}
