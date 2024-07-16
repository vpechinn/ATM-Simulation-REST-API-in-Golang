package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/accounts", createAccountHandler).Methods("POST")
	router.HandleFunc("/accounts/{id}/deposit", depositHandler).Methods("POST")
	router.HandleFunc("/accounts/{id}/withdraw", withdrawHandler).Methods("POST")
	router.HandleFunc("/accounts/{id}/balance", getBalanceHandler).Methods("GET")

	return router
}
