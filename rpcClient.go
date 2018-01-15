package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

const (
	accountAddress = "AcuhtcyXqRuwao2ayvqLVuQqh8YY34mor1"
	assetTypeNEO = "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
	assetTypeGAS = "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7"
)

	

type (
	NeoJSON struct {
		ID      int    `json:"id"`
		JSONRPC string `json:"jsonrpc"`
		Result  int `json:"result"`
	}
)

type (
	AccountMap struct {
		ID      int                    `json:"id"`
		JSONRPC string                 `json:"jsonrpc"`
		Result  map[string]interface{} `json:"result"`
	}
)


func GetInfo() (string, string){
	responseBlob, _ := http.Get("http://localhost:20332?jsonrpc=2.0&method=getaccountstate&params=['"+accountAddress+"']&id=1")
	buf, _ := ioutil.ReadAll(responseBlob.Body)
	
		type Balance struct {
			Asset  string `json:"asset"`
			Value string `json:"value"`
		}

		type Result struct {
			Balances []Balance `json:"balances"`
		}

		type Response struct {
			Result Result `json:"result"`
		}
	
		var response Response

	
	err :=json.Unmarshal(buf, &response)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	//return Neo and Gas amount
	return response.Result.Balances[0].Value, response.Result.Balances[1].Value
}


func Send(transaction Transaction) {
	//select between NEO and GAS ids
	assetType := ""
	switch at := transaction.AssetType; at {
	case "NEO":
		assetType = assetTypeNEO
	case "GAS":
		assetType = assetTypeGAS
	default:
		log.Fatal("error: asset type is unidentified")
	}
	
	rtcpRequest :="http://localhost:20332?jsonrpc=2.0&method=sendtoaddress&params=['"+assetType+"','"+transaction.Destination+"',"+transaction.Amount+"]&id=1"
	
	
	// for tests
	fmt.Println(rtcpRequest+"\n")
	response, _ := http.Get(rtcpRequest)
	buf, _ := ioutil.ReadAll(response.Body)
	// Unmarshall to map 
	mapConfig := make(map[string]interface{})
	json.Unmarshal(buf, &mapConfig)
	fmt.Printf("%+v\n", mapConfig["result"])
}