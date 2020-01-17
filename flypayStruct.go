package main

// Transaction encapsulates one transaction data
type Transaction struct {
	Amount         int    `json:"amount"`
	Currency       string `json:"currency"`
	StatusCode     int    `json:"statusCode"`
	OrderReference string `json:"orderReference"`
	TransactionID  string `json:"transactionId"`
}

// Params to filter transactions returned based on it
type Params struct {
	provider string
	jsonFile []byte
	status   string
	currency string
}
