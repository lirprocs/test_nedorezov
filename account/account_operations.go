package account

import (
	"fmt"
	"test_nedorezov/log"
)

type Operation struct {
	AccountID int
	Action    string
	Amount    float64
	Response  chan string
}

var operations chan Operation

func init() {
	operations = make(chan Operation)
	go processOperations()
}

func processOperations() {
	accounts := make(map[int]*Account)
	for op := range operations {
		switch op.Action {
		case "create":
			account := &Account{ID: op.AccountID}
			accounts[op.AccountID] = account
			op.Response <- fmt.Sprintf("Account %d created", op.AccountID)
		case "deposit":
			account := accounts[op.AccountID]
			if account != nil {
				account.Deposit(op.Amount)
				op.Response <- fmt.Sprintf("Deposited %f to account %d", op.Amount, op.AccountID)
			} else {
				op.Response <- fmt.Sprintf("Account %d not found", op.AccountID)
			}
		case "withdraw":
			account := accounts[op.AccountID]
			if account != nil {
				err := account.Withdraw(op.Amount)
				if err != nil {
					op.Response <- err.Error()
				} else {
					op.Response <- fmt.Sprintf("Withdrew %f from account %d", op.Amount, op.AccountID)
				}
			} else {
				op.Response <- fmt.Sprintf("Account %d not found", op.AccountID)
			}
		case "balance":
			account := accounts[op.AccountID]
			if account != nil {
				balance := account.GetBalance()
				op.Response <- fmt.Sprintf("Account %d balance: %f", op.AccountID, balance)
				op.Amount = balance
			} else {
				op.Response <- fmt.Sprintf("Account %d not found", op.AccountID)
			}
		}
		log.LogOperation(log.Operation{
			AccountID: op.AccountID,
			Action:    op.Action,
			Amount:    op.Amount,
		})
	}
}
