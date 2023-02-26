package repository

import (
	"Ewallet-infotecs/internal/model"
	"database/sql"
)

type TransactionHistory interface {
	Create(transaction model.Transaction) (int64, error)
	GetAll() ([]model.Transaction, error)
	GetById(transaction int) (model.Transaction, error)
	Update(transaction model.Transaction) error
	Delete(id int) error
}

type Wallet interface {
	Create(w model.Wallet) (int, error)
	GetAll() ([]model.Wallet, error)
	GetById(walletId int) (model.Wallet, error)
	GetByAddress(address string) (model.Wallet, error)
	Update(w model.Wallet) error
	Delete(id int) error
}

type Repository struct {
	TransactionHistory
	Wallet
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		TransactionHistory: NewTransactionHistorySqlite(db),
		Wallet:             NewWalletSqlite(db),
	}
}
