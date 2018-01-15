package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	//accountAddress = "AcuhtcyXqRuwao2ayvqLVuQqh8YY34mor1"
	assetTypeNEO = "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
	assetTypeGAS = "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7"
)

//from config.json
type Configuration struct {
	AccountAddress string
	Host           string
}



//returns current NEO and GAS balances
func GetInfo() []Balance {

	configuration := initConfig()

	//build url
	u, err := url.Parse(configuration.Host)
	if err != nil {
		log.Fatal(err)
	}

	paramsString := []string{"['", configuration.AccountAddress, "']"}
	params := strings.Join(paramsString, "")

	//u.Scheme = "http"
	//u.Host = "localhost:20332"
	q := u.Query()
	q.Set("jsonrpc", "2")
	q.Set("method", "getaccountstate")
	q.Set("id", "1")
	q.Set("params", params)
	u.RawQuery = q.Encode()

	//urlString := fmt.Sprint("http://localhost:20332?jsonrpc=2.0&method=getaccountstate&params=['%v']&id=1", configuration.AccountAddress)

	responseBlob, _ := http.Get(u.String())

	//
	fmt.Println(u.String())

	// responseBlob, _ := http.Get("http://localhost:20332?jsonrpc=2.0&method=getaccountstate&params=['"+accountAddress+"']&id=1")
	buf, _ := ioutil.ReadAll(responseBlob.Body)

	type Result struct {
		Balances []Balance `json:"balances"`
	}

	type Response struct {
		Result Result `json:"result"`
	}

	var response Response

	error := json.Unmarshal(buf, &response)
	if err != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(response)

	//return Neo and Gas amount

	return response.Result.Balances
}

//form and send API request to NEO node, transfer assets
func Send(transaction Transaction) {
	configuration := initConfig()

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

	//build url
	u, err := url.Parse(configuration.Host)
	if err != nil {
		log.Fatal(err)
	}

	paramsString := []string{"['", assetType, "',", transaction.Destination, "',", transaction.Amount, "]"}
	params := strings.Join(paramsString, "")

	//u.Scheme = "http"
	//u.Host = "localhost:20332"
	q := u.Query()
	q.Set("jsonrpc", "2")
	q.Set("method", "sendtoaddress")
	q.Set("id", "1")
	q.Set("params", params)
	u.RawQuery = q.Encode()
	fmt.Println(u)

	// rtcpRequest :="http://localhost:20332?jsonrpc=2.0&method=sendtoaddress&params=['"+assetType+"','"+transaction.Destination+"',"+transaction.Amount+"]&id=1"

	response, _ := http.Get(u.String())

	// for tests
	buf, _ := ioutil.ReadAll(response.Body)
	// Unmarshall to map
	mapConfig := make(map[string]interface{})
	json.Unmarshal(buf, &mapConfig)
	fmt.Printf("%+v\n", mapConfig["result"])
}

func initConfig() *Configuration {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &configuration
}
