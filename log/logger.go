package log

import (
	"fmt"
	//"test_nedorezov/account"
	"time"
)

type Operation struct {
	AccountID int
	Action    string
	Amount    float64
}

func LogOperation(op Operation) {
	fmt.Printf("%s - Account ID: %d, Action: %s, Amount: %f\n", time.Now().Format(time.RFC3339), op.AccountID, op.Action, op.Amount)
}
