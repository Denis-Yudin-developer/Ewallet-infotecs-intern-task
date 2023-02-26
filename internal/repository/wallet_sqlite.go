package repository

import (
	"database/sql"
	"log"
)
import "Ewallet-infotecs/internal/model"

type WalletSqlite struct {
	db *sql.DB
}

func (r *WalletSqlite) Create(wallet model.Wallet) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *WalletSqlite) GetAll() ([]model.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (r *WalletSqlite) GetById(walletId int) (model.Wallet, error) {
	log.Printf("Getting wallet by id: %d", walletId)
	row := r.db.QueryRow("select id, address, balance from wallet where id=?", walletId)

	w := model.Wallet{}
	if err := row.Scan(&w.ID, &w.Address, &w.Balance); err == sql.ErrNoRows {
		log.Printf("wallet not found")
		return model.Wallet{}, err
	}
	return w, nil
}

func (r *WalletSqlite) GetByAddress(address string) (model.Wallet, error) {
	log.Printf("Getting wallet by address: %s", address)

	row := r.db.QueryRow("select id, address, balance FROM wallet where address=?", address)

	w := model.Wallet{}
	var err error
	if err = row.Scan(&w.ID, &w.Address, &w.Balance); err == sql.ErrNoRows {
		log.Printf("wallet not found")
		return model.Wallet{}, err
	}
	return w, nil
}

func (r *WalletSqlite) Update(wallet model.Wallet) error {
	//TODO implement me
	_, err :=
		r.db.Exec("UPDATE wallet SET balance=? where id=?", wallet.Balance, wallet.ID)

	return err
}

func (r *WalletSqlite) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewWalletSqlite(db *sql.DB) *WalletSqlite {
	return &WalletSqlite{db: db}
}
