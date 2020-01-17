package main

import (
	"encoding/json"
)

// Transaction type A struct
type transactionA struct {
	Amount         int    `json:"amount"`
	Currency       string `json:"currency"`
	StatusCode     int    `json:"statusCode"`
	OrderReference string `json:"orderReference"`
	TransactionID  string `json:"transactionId"`
}

// Transactions list type A struct
type transactionsContainerA struct {
	Transactions []transactionA `json:"transactions"`
}

// NewTransactionsFromA initializes transactions from JSON with format A
func NewTransactionsFromA(params Params) (transactions []*Transaction) {

	statusCode := 0

	if len(params.status) > 0 {
		if params.status == "authorised" {
			statusCode = 1
		} else if params.status == "decline" {
			statusCode = 2
		} else if params.status == "refunded" {
			statusCode = 3
		} else {
			statusCode = 0
		}
	}

	currency := "any"

	if len(params.currency) > 0 {
		currency = params.currency
	}

	var container = &transactionsContainerA{}

	json.Unmarshal(params.jsonFile, container)

	transactions = make([]*Transaction, len(container.Transactions))

	for index, transaction := range container.Transactions {
		if statusCode == 0 && currency == "any" {
			transactions[index] = &Transaction{
				Amount:         transaction.Amount,
				Currency:       transaction.Currency,
				OrderReference: transaction.OrderReference,
				StatusCode:     transaction.StatusCode,
				TransactionID:  transaction.TransactionID,
			}
		} else if statusCode > 0 && transaction.StatusCode == statusCode && currency == "any" {
			transactions[index] = &Transaction{
				Amount:         transaction.Amount,
				Currency:       transaction.Currency,
				OrderReference: transaction.OrderReference,
				StatusCode:     transaction.StatusCode,
				TransactionID:  transaction.TransactionID,
			}
		} else if statusCode > 0 && transaction.StatusCode == statusCode && currency != "any" {
			if transaction.Currency == currency {
				transactions[index] = &Transaction{
					Amount:         transaction.Amount,
					Currency:       transaction.Currency,
					OrderReference: transaction.OrderReference,
					StatusCode:     transaction.StatusCode,
					TransactionID:  transaction.TransactionID,
				}
			}
		} else if statusCode == 0 && currency != "any" {
			if transaction.Currency == currency {
				transactions[index] = &Transaction{
					Amount:         transaction.Amount,
					Currency:       transaction.Currency,
					OrderReference: transaction.OrderReference,
					StatusCode:     transaction.StatusCode,
					TransactionID:  transaction.TransactionID,
				}
			}
		}
	}

	return
}
