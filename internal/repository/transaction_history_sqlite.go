package repository

import (
	"Ewallet-infotecs/internal/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type TransactionHistorySqlite struct {
	db *sql.DB
}

func (t TransactionHistorySqlite) Create(transaction model.Transaction) (int64, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	timeNow := time.Now()

	query := fmt.Sprintf("INSERT INTO"+
		" transaction_history(from_id, to_id, transfer_amount, transfer_time) "+
		"values ((SELECT id from wallet where address=\"%s\"), (SELECT id from wallet where address=\"%s\"), %f, \"%s\")",
		transaction.FromAddress,
		transaction.ToAddress,
		transaction.TransferAmount,
		timeNow.Format(time.RFC3339))

	row, err :=
		tx.Exec(query)
	itemId, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	log.Printf("item id is %d", itemId)

	return itemId, tx.Commit()
}

func (t TransactionHistorySqlite) GetAll() ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)
	rows, err := t.db.Query("select tr.id, tr.transfer_amount, tr.transfer_time, w.address, w1.address from transaction_history tr INNER JOIN wallet w on w.id = tr.from_id INNER JOIN wallet w1 on w1.id = tr.to_id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var transaction model.Transaction

		err = rows.Scan(&transaction.ID,
			&transaction.TransferAmount,
			&transaction.TransferTime,
			&transaction.FromAddress,
			&transaction.ToAddress)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t TransactionHistorySqlite) GetById(transaction int) (model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t TransactionHistorySqlite) Update(transaction model.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t TransactionHistorySqlite) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewTransactionHistorySqlite(db *sql.DB) *TransactionHistorySqlite {
	return &TransactionHistorySqlite{db: db}
}
