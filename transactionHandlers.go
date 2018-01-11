package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Transaction : Plain Old Transaction
type Transaction struct {
	//AssetAddress string `json:"asset"`
	Destination string `json:"destination"`
	Amount string `json:"amount"`
	AssetType string `json:"type"`
}

type Account struct {
	NeoAsset string `json:"neo"`
	GasAsset string `json:"gas"`
}


var Transactions []Transaction

func getAccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "Transactions" variable to json
	neoValue, gasValue := getInfo();
	account := Account{neoValue, gasValue}
	AccountBytes, err := json.Marshal(account)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 
	w.Write(AccountBytes)
}

func transferHandler(w http.ResponseWriter, r *http.Request) {
	Transaction := Transaction{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Transaction.AssetAddress = r.Form.Get("asset")
	Transaction.Destination = r.Form.Get("destination")
	Transaction.Amount = r.Form.Get("amount")
	Transaction.AssetType = r.Form.Get("type")

	Transactions = append(Transactions, Transaction)
	fmt.Println(Transaction);
	send(Transaction)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}