package account

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateAccountHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccountHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if response == "" {
		t.Errorf("Expected non-empty response")
	}
}

func TestDepositHandler(t *testing.T) {
	accountID := generateAccountID()
	CreateAccount(accountID) // Assuming CreateAccount is a helper function to create an account

	var jsonStr = []byte(`{"amount": 100}`)
	req, err := http.NewRequest("POST", "/accounts/"+strconv.Itoa(accountID)+"/deposit", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DepositHandler)

	r := mux.NewRouter()
	r.HandleFunc("/accounts/{id}/deposit", handler).Methods("POST")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if response == "" {
		t.Errorf("Expected non-empty response")
	}
}

func TestWithdrawHandler(t *testing.T) {
	accountID := generateAccountID()
	CreateAccount(accountID)
	Deposit(accountID, 200)

	var jsonStr = []byte(`{"amount": 50}`)
	req, err := http.NewRequest("POST", "/accounts/"+strconv.Itoa(accountID)+"/withdraw", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WithdrawHandler)

	r := mux.NewRouter()
	r.HandleFunc("/accounts/{id}/withdraw", handler).Methods("POST")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if response == "" {
		t.Errorf("Expected non-empty response")
	}
}

func TestBalanceHandler(t *testing.T) {
	accountID := generateAccountID()
	CreateAccount(accountID)
	Deposit(accountID, 300)

	req, err := http.NewRequest("GET", "/accounts/"+strconv.Itoa(accountID)+"/balance", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BalanceHandler)

	r := mux.NewRouter()
	r.HandleFunc("/accounts/{id}/balance", handler).Methods("GET")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if response == "" {
		t.Errorf("Expected non-empty response")
	}
}

// Helper functions
func CreateAccount(id int) {
	response := make(chan string)
	operations <- Operation{AccountID: id, Action: "create", Response: response}
	<-response
}

func Deposit(id int, amount float64) {
	response := make(chan string)
	operations <- Operation{AccountID: id, Action: "deposit", Amount: amount, Response: response}
	<-response
}
