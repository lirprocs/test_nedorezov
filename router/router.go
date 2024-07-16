package router

import (
	"github.com/gorilla/mux"
	"test_nedorezov/account"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", account.CreateAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/deposit", account.DepositHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/withdraw", account.WithdrawHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", account.BalanceHandler).Methods("GET")
	return r
}
