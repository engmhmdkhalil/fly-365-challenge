package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Creating new router
	router := mux.NewRouter()
	// Adding prefix
	api := router.PathPrefix("/api/payment").Subrouter()

	api.HandleFunc("/transactaion", GetTransactions).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
