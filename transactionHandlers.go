package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Transaction : Plain Old Transaction
type Transaction struct {
	Destination string `json:"destination"`
	Amount string `json:"amount"`
	AssetType string `json:"type"`
}

type Account struct {
	NeoAsset string `json:"neo"`
	GasAsset string `json:"gas"`
}


var Transactions []Transaction

func AccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	neoValue, gasValue := GetInfo();
	account := Account{neoValue, gasValue}
	AccountBytes, err := json.Marshal(account)

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
	fmt.Println(Transaction);
	Send(Transaction)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}