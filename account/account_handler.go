package account

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	id := generateAccountID()
	response := make(chan string)
	operations <- Operation{AccountID: id, Action: "create", Response: response}
	res := <-response
	json.NewEncoder(w).Encode(res)
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var amount struct {
		Amount float64 `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&amount)
	response := make(chan string)
	operations <- Operation{AccountID: id, Action: "deposit", Amount: amount.Amount, Response: response}
	res := <-response
	json.NewEncoder(w).Encode(res)
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var amount struct {
		Amount float64 `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&amount)
	response := make(chan string)
	operations <- Operation{AccountID: id, Action: "withdraw", Amount: amount.Amount, Response: response}
	res := <-response
	json.NewEncoder(w).Encode(res)
}

func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	response := make(chan string)
	operations <- Operation{AccountID: id, Action: "balance", Response: response}
	res := <-response
	json.NewEncoder(w).Encode(res)
}

func generateAccountID() int {
	return int(time.Now().UnixNano())
}
