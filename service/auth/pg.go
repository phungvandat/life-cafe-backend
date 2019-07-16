package auth

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"

	"github.com/phungvandat/life-cafe-backend/model/domain"
	"github.com/phungvandat/life-cafe-backend/util/contextkey"
)

// pgService implmenter for auth service in postgres
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

func (s *pgService) AuthenticateUser(ctx context.Context) error {
	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)
	if !check {
		return NotLoggedInError
	}
	userID, err := domain.UUIDFromString(ctxUserID)
	if err != nil {
		return err
	}
	user := &domain.User{Model: domain.Model{
		ID: userID,
	}}

	err = s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return AccountNotFoundError
	}

	if user.Active == false {
		return AccountIsLockedError
	}

	return nil
}

func (s *pgService) AuthenticateAdmin(ctx context.Context) error {
	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)
	if !check {
		return NotLoggedInError
	}
	userID, err := domain.UUIDFromString(ctxUserID)
	if err != nil {
		return err
	}
	user := &domain.User{Model: domain.Model{
		ID: userID,
	}}

	err = s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return AccountNotFoundError
	}

	if user.Active == false {
		return AccountIsLockedError
	}

	if user.Role != "admin" {
		return AccessDeniedError
	}

	return nil
}
