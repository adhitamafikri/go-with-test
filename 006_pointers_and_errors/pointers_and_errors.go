package pointers_and_errors

import (
	"errors"
)

var ErrInitialBalanceTooSmall = errors.New("cannot create wallet with initial balance lesser than 10.0")
var ErrInsufficientBalance = errors.New("cannot withdraw. insufficient balance")

type Wallet struct {
	balance float64
}

func NewWallet(balance float64) (*Wallet, error) {
	if balance < 10.0 {
		return nil, ErrInitialBalanceTooSmall
	}
	return &Wallet{balance: balance}, nil
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Withdraw(amount float64) error {
	if w.Balance() < amount {
		return ErrInsufficientBalance
	}

	newBalance := w.Balance() - amount
	w.balance = newBalance
	return nil
}
