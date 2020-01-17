package main

import (
	"encoding/json"
)

// Transaction type B struct
type transactionB struct {
	Value               int    `json:"value"`
	TransactionCurrency string `json:"transactionCurrency"`
	StatusCode          int    `json:"statusCode"`
	OrderInfo           string `json:"orderInfo"`
	PaymentID           string `json:"paymentId"`
}

// Transactions list type A struct
type transactionsContainerB struct {
	Transactions []transactionB `json:"transactions"`
}

// NewTransactionsFromB initializes transactions from JSON with format B
func NewTransactionsFromB(params Params) (transactions []*Transaction) {

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

	var container = &transactionsContainerB{}

	json.Unmarshal(params.jsonFile, container)

	transactions = make([]*Transaction, len(container.Transactions))

	for index, transaction := range container.Transactions {
		if statusCode == 0 && currency == "any" {
			transactions[index] = &Transaction{
				Amount:         transaction.Value,
				Currency:       transaction.TransactionCurrency,
				StatusCode:     transaction.StatusCode / 100,
				OrderReference: transaction.OrderInfo,
				TransactionID:  transaction.PaymentID,
			}
		} else if statusCode > 0 && transaction.StatusCode/100 == statusCode && currency == "any" {
			transactions[index] = &Transaction{
				Amount:         transaction.Value,
				Currency:       transaction.TransactionCurrency,
				StatusCode:     transaction.StatusCode / 100,
				OrderReference: transaction.OrderInfo,
				TransactionID:  transaction.PaymentID,
			}
		} else if statusCode > 0 && transaction.StatusCode/100 == statusCode && currency != "any" {
			if transaction.TransactionCurrency == currency {
				transactions[index] = &Transaction{
					Amount:         transaction.Value,
					Currency:       transaction.TransactionCurrency,
					StatusCode:     transaction.StatusCode / 100,
					OrderReference: transaction.OrderInfo,
					TransactionID:  transaction.PaymentID,
				}
			}
		} else if statusCode == 0 && currency != "any" {
			if transaction.TransactionCurrency == currency {
				transactions[index] = &Transaction{
					Amount:         transaction.Value,
					Currency:       transaction.TransactionCurrency,
					StatusCode:     transaction.StatusCode / 100,
					OrderReference: transaction.OrderInfo,
					TransactionID:  transaction.PaymentID,
				}
			}
		}
	}

	return
}
