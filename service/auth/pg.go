package auth

import (
	"context"

	"github.com/jinzhu/gorm"

	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	"github.com/phungvandat/life-cafe-backend/util/contextkey"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
)

// pgService implmenter for auth service in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService new pg service
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func (s *pgService) AuthenticateUser(ctx context.Context) error {
	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)
	if !check {
		return errors.NotLoggedInError
	}
	userID, err := pgModel.UUIDFromString(ctxUserID)
	if err != nil {
		return err
	}
	user := &pgModel.User{Model: pgModel.Model{
		ID: userID,
	}}

	err = s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return errors.AccountNotFoundError
	}

	if user.Active == false {
		return errors.AccountIsLockedError
	}

	return nil
}

func (s *pgService) AuthenticateAdmin(ctx context.Context) error {
	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)
	if !check {
		return errors.NotLoggedInError
	}
	userID, err := pgModel.UUIDFromString(ctxUserID)
	if err != nil {
		return err
	}
	user := &pgModel.User{Model: pgModel.Model{
		ID: userID,
	}}

	err = s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return errors.AccountNotFoundError
	}

	if user.Active == false {
		return errors.AccountIsLockedError
	}

	if user.Role != "admin" {
		return errors.AccessDeniedError
	}

	return nil
}
