package service

import (
	"Ewallet-infotecs/internal/model"
	"Ewallet-infotecs/internal/repository"
)

type WalletService struct {
	repo repository.Wallet
}

func (w *WalletService) Create(wallet model.Wallet) (int, error) {
	panic("implement me")
}

func (w *WalletService) GetAll() ([]model.Wallet, error) {
	panic("implement me")
}

func (w *WalletService) GetById(wallet int) (model.Wallet, error) {
	panic("implement me")
}

func (w *WalletService) GetByAddress(address string) (model.Wallet, error) {
	return w.repo.GetByAddress(address)
}

func (w *WalletService) Update(wallet model.Wallet) error {
	panic("implement me")
}

func (w *WalletService) Delete(id int) error {
	panic("implement me")
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}
