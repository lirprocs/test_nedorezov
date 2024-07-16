package account

import (
	"fmt"
	"sync"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      int
	Balance float64
	Mu      sync.Mutex
}

func (a *Account) Deposit(amount float64) error {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	if amount > a.Balance {
		return fmt.Errorf("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	return a.Balance
}
