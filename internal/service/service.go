package service

import (
	"Ewallet-infotecs/internal/model"
	"Ewallet-infotecs/internal/repository"
)

type TransactionHistory interface {
	Create(transaction model.Transaction) (int64, error)
	GetAll() ([]model.Transaction, error)
	GetById(transaction int) (model.Transaction, error)
	Update(transaction model.Transaction) error
	LastNTransaction(transactionsNumber int) ([]model.Transaction, error)
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

type Service struct {
	TransactionHistory
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TransactionHistory: NewTransactionHistoryService(repos.TransactionHistory, repos.Wallet),
		Wallet:             NewWalletService(repos.Wallet),
	}
}
