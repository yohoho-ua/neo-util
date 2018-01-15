package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Transaction : Plain Old Transaction
type Transaction struct {
	Destination string `json:"destination"`
	Amount      string `json:"amount"`
	AssetType   string `json:"type"`
}

type BalancesArray struct {
    Balances [] Balance `json:"balances"`
}

type (
	Balance struct {
		Asset string `json:"asset"`
		Value string `json:"value"`              
	}
)


var Transactions []Transaction

func AccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	
	balances  := GetInfo()

	fmt.Println(balances)
	fmt.Println(len(balances))
	
	AccountBytes, err := json.Marshal(balances)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(AccountBytes)
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	Transaction := Transaction{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Transaction.Destination = r.Form.Get("destination")
	Transaction.Amount = r.Form.Get("amount")
	Transaction.AssetType = r.Form.Get("type")

	Transactions = append(Transactions, Transaction)
	fmt.Println(Transaction)
	Send(Transaction)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
