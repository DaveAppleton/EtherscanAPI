package etherscanAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TxListItem struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	IsError           string `json:"isError"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
}

// TxListRec - result from Transaction Calls
//   Status  - OK / NOTOK
//   Message - error if Status NOTOK
//   Result  - TxListItem
type TxListRec struct {
	Status  string
	Message string
	Result  []TxListItem
}

// TransactionsByAddress : Get a list of 'Normal' Transactions By Address
// [Optional Parameters] startblock: starting blockNo to retrieve results, endblock: ending blockNo to retrieve results
// ([BETA] Returned 'isError' values: 0=No Error, 1=Got Error)
// (Returns up to a maximum of the last 10000 transactions only)
// http://api.etherscan.io/api?module=account&action=txlist&address=0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae&startblock=0&endblock=99999999&sort=asc&apikey=YourApiKeyToken
// (To get paginated results use page=<page number> and offset=<max records to return>)
// https://api.etherscan.io/api?module=account&action=txlist&address=0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae&startblock=0&endblock=99999999&page=1&offset=10&sort=asc&apikey=YourApiKeyToken
func (a *API) TransactionsByAddress(addr string) (tr TxListRec) {
	//var tr txListRec
	call := "http://api.etherscan.io/api?module=account&action=txlist&startblock=0&endblock=99999999&sort=asc" + "&address=" + addr + "&tag=latest&apikey=" + a.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return TxListRec{Status: "NOTOK", Message: err.Error()}
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), 400)
		return TxListRec{Status: "NOTOK", Message: err.Error()}
	}
	return
}

// Get a list of 'Internal' Transactions by Address
// [Optional Parameters] startblock: starting blockNo to retrieve results, endblock: ending blockNo to retrieve results
// (Returned 'isError' values: 0=No Error, 1=Got Error)
// (Returns up to a maximum of the last 10000 transactions only)

// http://api.etherscan.io/api?module=account&action=txlistinternal&address=0x2c1ba59d6f58433fb1eaee7d20b26ed83bda51a3&startblock=0&endblock=2702578&sort=asc&apikey=YourApiKeyToken

// (To get paginated results use page=<page number> and offset=<max records to return>)
// https://api.etherscan.io/api?module=account&action=txlistinternal&address=0x2c1ba59d6f58433fb1eaee7d20b26ed83bda51a3&startblock=0&endblock=2702578&page=1&offset=10&sort=asc&apikey=YourApiKeyToken

func (a *API) internalTransactionsByAddress(addr string) (tr TxListRec) {
	call := "http://api.etherscan.io/api?module=account&action=txlistinternal&address=" + addr + "&startblock=0&endblock=2702578&sort=asc&apikey=" + a.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return TxListRec{Status: "NOTOK", Message: err.Error()}
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), 400)
		return TxListRec{Status: "NOTOK", Message: err.Error()}
	}
	return
}

// Get "Internal Transactions" by Transaction Hash
// https://api.etherscan.io/api?module=account&action=txlistinternal&txhash=0x40eb908387324f2b575b4879cd9d7188f69c8fc9d87c901b9e2daaea4b442170&apikey=YourApiKeyToken
