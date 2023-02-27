package model

import (
	"time"
)

type Transaction struct {
	ID             int64     `json:"id"`
	FromAddress    string    `json:"from"`
	ToAddress      string    `json:"to"`
	TransferAmount float64   `json:"transfer_amount"`
	TransferTime   time.Time `json:"transfer_time"`
}
