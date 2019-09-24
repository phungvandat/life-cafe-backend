package sagas

import (
	"github.com/jinzhu/gorm"
)

// SagasService interface
type SagasService interface {
	NewTransaction(string, *gorm.DB)
	CommitTransaction(string) error
	RollbackTransaction(string) error
	CommitAllTransaction() error
	RollbackAllTransaction() error
}

// TransactionRollback struct
type TransactionRollback struct {
	transaction map[string]*gorm.DB
}

// NewSagasService func
func NewSagasService() SagasService {
	transaction := make(map[string]*gorm.DB)
	return &TransactionRollback{
		transaction: transaction,
	}
}

// NewTransaction func
func (tr *TransactionRollback) NewTransaction(transactionID string, tx *gorm.DB) {
	tr.transaction[transactionID] = tx
}

// CommitTransaction func
func (tr *TransactionRollback) CommitTransaction(transactionID string) error {
	tx := tr.transaction[transactionID]
	if tx == nil {
		return TransactionRollbackeNotExistError
	}
	delete(tr.transaction, transactionID)
	return tx.Commit().Error
}

// RollbackTransaction func
func (tr *TransactionRollback) RollbackTransaction(transactionID string) error {
	tx := tr.transaction[transactionID]
	if tx == nil {
		return TransactionRollbackeNotExistError
	}
	delete(tr.transaction, transactionID)
	return tx.Rollback().Error
}

// CommitAllTransaction func
func (tr *TransactionRollback) CommitAllTransaction() error {
	for key, tx := range tr.transaction {
		delete(tr.transaction, key)
		err := tx.Commit().Error
		if err != nil {
			return err
		}
	}
	return nil
}

// RollbackAllTransaction func
func (tr *TransactionRollback) RollbackAllTransaction() error {
	for key, tx := range tr.transaction {
		delete(tr.transaction, key)
		err := tx.Rollback().Error
		if err != nil {
			return err
		}
	}
	return nil
}
