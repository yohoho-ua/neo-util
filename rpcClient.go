package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	accountAddress = "AcuhtcyXqRuwao2ayvqLVuQqh8YY34mor1"
)

	

type (
	// String represents the JSON schema of a response from a NEO node, where the expected
	// result is a string.
	NeoJSON struct {
		ID      int    `json:"id"`
		JSONRPC string `json:"jsonrpc"`
		Result  int `json:"result"`
	}
)

type (
	// String represents the JSON schema of a response from a NEO node, where the expected
	// result is a string.
	AccountMap struct {
		ID      int                    `json:"id"`
		JSONRPC string                 `json:"jsonrpc"`
		Result  map[string]interface{} `json:"result"`
	}
)

// error checking omitted for brevity
// func main() {
// var neoResponse NeoJSON
// response, _ := http.Get("http://localhost:20332?jsonrpc=2.0&method=getblockcount&params=[]&id=1")
// buf, _ := ioutil.ReadAll(response.Body)
// json.Unmarshal(buf, &neoResponse)
// fmt.Println(neoResponse.Result)

// }

// func main() {
// 	var f interface{}
// 	//var m map[string]int
// //s := `{"a":1, "b":2, "x":1, "y":1}`
// response, _ := http.Get("http://localhost:20332?jsonrpc=2.0&method=getblockcount&params=[]&id=1")
// buf, _ := ioutil.ReadAll(response.Body)

// if err := json.Unmarshal(buf, &f); err != nil {
// 	panic(err)
// }

// fmt.Printf("%+v", f)
// }


func getInfo() (string, string){
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


func send(transaction Transaction) {
	//select between NEO and GAS ids
	assetType := "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b" //NEO
	if transaction.AssetType != "NEO" {
		assetType = "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7" //GAS
}	
	rtcpRequest :="http://localhost:20332?jsonrpc=2.0&method=sendtoaddress&params=['"+assetType+"','"+transaction.Destination+"',"+transaction.Amount+"]&id=1"
	fmt.Println(rtcpRequest+"\n")
	response, _ := http.Get(rtcpRequest)
	// response, _ := http.Get("http://localhost:20332?jsonrpc=2.0&method=sendtoaddress&params=['0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7','AHcwixWbxYbfPirmFXA7dX13bu6Po75UCV',33]&id=1")
	buf, _ := ioutil.ReadAll(response.Body)
	// Unmarshall to map
	mapConfig := make(map[string]interface{})
	json.Unmarshal(buf, &mapConfig)
	fmt.Printf("%+v\n", mapConfig["result"])
}