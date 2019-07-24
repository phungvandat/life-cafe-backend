package photo

import (
	"context"

	"github.com/jinzhu/gorm"
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/helper"
)

// pgService implmenter for photo serivce in postgres
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

func (s *pgService) CreatePhoto(ctx context.Context, req requestModel.CreatePhotoRequest) (*responseModel.CreatePhotoResponse, error) {
	tx := s.db.Begin()
	transactionID := (pgModel.NewUUID()).String()
	s.spRollback.NewTransaction(transactionID, tx)

	res := &responseModel.CreatePhotoResponse{
		TransactionID: &transactionID,
	}

	photo := &pgModel.Photo{
		URL: req.URL,
	}

	if req.ProductID != "" {
		productIDUUID, _ := pgModel.UUIDFromString(req.ProductID)
		photo.ProductID = &productIDUUID
	}

	if req.PhotoID != "" {
		photoIDUUID, _ := pgModel.UUIDFromString(req.PhotoID)
		photo.ID = photoIDUUID
	}
	err := tx.Create(photo).Error

	if err != nil {
		return res, err
	}
	res.Photo = photo

	return res, nil
}

func (s *pgService) GetPhotos(ctx context.Context, req requestModel.GetPhotosRequest) (*responseModel.GetPhotosResponse, error) {
	res := &responseModel.GetPhotosResponse{}
	photos := []*pgModel.Photo{}

	skip := req.Skip
	limit := req.Limit
	if req.Skip == "" {
		skip = "-1"
	}

	if req.Limit == "" {
		limit = "-1"
	}

	stringQuery := ""
	if req.ProductID != "" {
		stringQuery += "product_id = '" + req.ProductID + "'"
	}

	err := s.db.
		Limit(limit).
		Offset(skip).
		Table("photos").
		Where(stringQuery).
		Scan(&photos).Error

	if err != nil {
		return res, err
	}

	res.Photos = photos

	return res, nil
}

func (s *pgService) RemovePhoto(ctx context.Context, req requestModel.RemovePhotoRequest) (*responseModel.RemovePhotoResponse, error) {
	tx := s.db.Begin()
	transactionID := (pgModel.NewUUID()).String()
	s.spRollback.NewTransaction(transactionID, tx)
	res := &responseModel.RemovePhotoResponse{
		TransactionID: &transactionID,
	}
	stringQuery := ""
	if req.ParamPhotoID != "" {
		stringQuery += "id = '" + req.ParamPhotoID + "'"
	}

	if req.ProductID != "" {
		productQuery := "product_id = '" + req.ProductID + "'"
		if stringQuery != "" {
			stringQuery += " AND " + productQuery
		} else {
			stringQuery += productQuery
		}
	}
	err := tx.Unscoped().Where(stringQuery).Delete(&pgModel.Photo{}).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *pgService) GetPhoto(ctx context.Context, req requestModel.GetPhotoRequest) (*responseModel.GetPhotoResponse, error) {
	res := &responseModel.GetPhotoResponse{}

	photoIDUUID, _ := pgModel.UUIDFromString(req.ParamPhotoID)

	photo := &pgModel.Photo{
		Model: pgModel.Model{
			ID: photoIDUUID,
		},
	}
	err := s.db.Find(photo, photo).Error

	if err != nil {
		return res, err
	}

	res.Photo = photo

	return res, nil
}
