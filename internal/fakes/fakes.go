package fakes

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var walletSymbols = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

const walletLenght = 64

func FakeWallet() string {
	b := make([]rune, walletLenght)
	for i := 0; i < walletLenght; i++ {
		b[i] = walletSymbols[rand.Intn(len(walletSymbols))]
	}
	return string(b)
}
