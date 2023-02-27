package repository

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)
import "Ewallet-infotecs/internal/model"

type WalletSqlite struct {
	db *sql.DB
}

func (r *WalletSqlite) Create(wallet model.Wallet) (int, error) {
	panic("implement me")
}

func (r *WalletSqlite) GetAll() ([]model.Wallet, error) {
	panic("implement me")
}

func (r *WalletSqlite) GetById(walletId int) (model.Wallet, error) {
	logrus.Print("Поиск кошелька с id %d", walletId)
	row := r.db.QueryRow("select id, address, balance from wallet where id=?", walletId)

	w := model.Wallet{}
	err := row.Scan(&w.ID, &w.Address, &w.Balance)
	if err != nil {
		return w, err
	}

	return w, nil
}

func (r *WalletSqlite) GetByAddress(address string) (model.Wallet, error) {
	logrus.Printf("Поиск кошелька по адресу %s", address)
	row := r.db.QueryRow("select id, address, balance FROM wallet where address=?", address)

	w := model.Wallet{}
	err := row.Scan(&w.ID, &w.Address, &w.Balance)
	if err != nil {
		return model.Wallet{}, err
	}
	return w, nil
}

func (r *WalletSqlite) Update(wallet model.Wallet) error {
	logrus.Printf("Обновление баланса кошелька с id: %d", wallet.ID)
	_, err :=
		r.db.Exec("UPDATE wallet SET balance=? where id=?", wallet.Balance, wallet.ID)

	return err
}

func (r *WalletSqlite) Delete(id int) error {
	panic("implement me")
}

func NewWalletSqlite(db *sql.DB) *WalletSqlite {
	return &WalletSqlite{db: db}
}
