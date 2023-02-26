package service

import (
	"Ewallet-infotecs/internal/model"
	"Ewallet-infotecs/internal/repository"
	"fmt"
	"github.com/pkg/errors"
)

type TransactionHistoryService struct {
	repo       repository.TransactionHistory
	walletRepo repository.Wallet
}

func (t TransactionHistoryService) LastNTransaction(transactionNumbers int) ([]model.Transaction, error) {
	allTransactions, _ := t.repo.GetAll()
	if len(allTransactions) < transactionNumbers || transactionNumbers < 0 {
		err := errors.New("Число необходимых для получения записей слишком велико")
		return nil, err
	}
	return allTransactions[len(allTransactions)-transactionNumbers:], nil
}

func (t TransactionHistoryService) Create(transaction model.Transaction) (int64, error) {
	if transaction.FromAddress == transaction.ToAddress {
		errorMessage := "Невозможно совершить транзакцию на свой адрес"
		return 0, errors.New(errorMessage)
	}

	from, err := t.walletRepo.GetByAddress(transaction.FromAddress)
	if err != nil {
		return 0, err
	}

	to, err := t.walletRepo.GetByAddress(transaction.ToAddress)
	if err != nil {
		return 0, err
	}

	if from.Balance < transaction.TransferAmount {
		errorMessage := fmt.Sprintf("НА КОШЕЛЬКЕ %s НЕДОСТАТОЧНО ДЕНЕГ.", from.Address)
		return 0, errors.New(errorMessage)
	}

	id, err := t.repo.Create(transaction)
	if err != nil {
		return 0, err
	}

	from.Balance -= transaction.TransferAmount

	err = t.walletRepo.Update(from)
	if err != nil {
		return 0, nil
	}

	to.Balance += transaction.TransferAmount
	err = t.walletRepo.Update(to)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t TransactionHistoryService) GetAll() ([]model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t TransactionHistoryService) GetById(transaction int) (model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t TransactionHistoryService) Update(transaction model.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t TransactionHistoryService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewTransactionHistoryService(repo repository.TransactionHistory, walletRepo repository.Wallet) *TransactionHistoryService {
	return &TransactionHistoryService{repo: repo, walletRepo: walletRepo}
}
