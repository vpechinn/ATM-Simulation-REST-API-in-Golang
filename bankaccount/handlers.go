package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	accounts = make(map[string]*Account)
	mu       sync.Mutex
)

type AccountRequest struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req AccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exist := accounts[req.ID]; exist {
		http.Error(w, "account already exist", http.StatusBadRequest)
		return
	}

	accounts[req.ID] = &Account{ID: req.ID}
	logOperation(req.ID, "CreateAccount", 0)
	w.WriteHeader(http.StatusCreated)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	var req AccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	account, exists := accounts[req.ID]
	mu.Unlock()

	if !exists {
		http.Error(w, "account not found", http.StatusNotFound)
	}

	go func() {
		if err := account.Deposit(req.Amount); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}()
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	var req AccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	account, exists := accounts[req.ID]
	mu.Unlock()

	if !exists {
		http.Error(w, "account not found", http.StatusNotFound)
		return
	}

	go func() {
		if err := account.Withdraw(req.Amount); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}()
}

func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	mu.Lock()
	account, exists := accounts[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "account not found", http.StatusNotFound)
		return
	}

	go func() {
		balance := account.GetBalance()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
	}()
}
