package main

import (
	"encoding/json"
	"net/http"
)

// GetTransactions , return all transactions combined with applying filters if any
func GetTransactions(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	var transactions []*Transaction

	if len(queryParams) > 0 {

		transactions = GetAllTransactions(Params{provider: queryParams.Get("provider"), status: queryParams.Get("status"), currency: queryParams.Get("currency")})
	} else {
		transactions = GetAllTransactions(Params{})
	}

	response, _ := json.Marshal(transactions)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
