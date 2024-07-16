package main

import (
	"errors"
	"sync"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      string
	Balance float64
	mu      sync.Mutex
}

func (a *Account) Deposit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Lock()

	if amount <= 0 {
		return errors.New("where is my money")
	}
	a.Balance += amount
	logOperation(a.ID, "Deposit", amount)
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	if a.Balance < amount {
		return errors.New("summa previshaet balance")
	}
	logOperation(a.ID, "Withdraw", amount)
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()

	logOperation(a.ID, "GetBalance", 0)
	return a.Balance
}
