package model

type Wallet struct {
	ID      int64   `json:"id"`
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}
