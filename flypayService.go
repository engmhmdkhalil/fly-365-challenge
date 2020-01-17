package main

import "io/ioutil"

// GetAllTransactions ... return all transactions from both files A & B with filters if any
func GetAllTransactions(params Params) (transactions []*Transaction) {

	// Get fly pay A transactions only
	if params.provider == "flypayA" {

		transactions = getFlyPayA(params)

		// Get fly pay B transactions only
	} else if params.provider == "flypayB" {

		transactions = getFlyPayB(params)

		// Get transactions from both files with applying filters if any
	} else if len(params.provider) == 0 {

		transactions = append(getFlyPayA(params), getFlyPayB(params)...)
	}

	return
}

// read flypayA and get transactions from it
func getFlyPayA(params Params) (transactions []*Transaction) {

	// Getting transactions from file A
	flypayA, _ := ioutil.ReadFile("flypayA.json")
	transactionsA := NewTransactionsFromA(Params{jsonFile: flypayA, status: params.status, currency: params.currency})

	for _, transaction := range transactionsA {
		if transaction != nil {
			transactions = append(transactions, transaction)
		}
	}

	return
}

// read flypayB and get transactions from it
func getFlyPayB(params Params) (transactions []*Transaction) {

	// Getting transactions from file B
	flypayB, _ := ioutil.ReadFile("flypayB.json")
	transactionsB := NewTransactionsFromB(Params{jsonFile: flypayB, status: params.status, currency: params.currency})

	for _, transaction := range transactionsB {
		if transaction != nil {
			transactions = append(transactions, transaction)
		}
	}

	return
}
