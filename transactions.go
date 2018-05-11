package etherscanAPI

/******************************************************************
Copyright 2017 David Appleton

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
******************************************************************/
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
	ContractAddress   string `json:"contractAddress"`
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
	// //var tr txListRec
	// call := "http://api.etherscan.io/api?module=account&action=txlist&startblock=0&endblock=99999999&sort=asc" + "&address=" + addr + "&tag=latest&apikey=" + a.apiKey
	// fmt.Println(call)
	// resp, err := http.Get(call)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return TxListRec{Status: "NOTOK", Message: err.Error()}
	// }
	// err = json.NewDecoder(resp.Body).Decode(&tr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	//http.Error(w, err.Error(), 400)
	// 	return TxListRec{Status: "NOTOK", Message: err.Error()}
	// }
	// return
	return a.AccountTransactions("txlist", "address", addr, "0", "99999999", "asc")
}

// TransactionsByAddressFromTo - return transactions in specified address range. empty string to use default
func (a *API) TransactionsByAddressFromTo(addr string, from string, to string) (tr TxListRec) {
	if len(from) == 0 {
		from = "0"
	}
	if len(to) == 0 {
		to = "99999999"
	}
	return a.AccountTransactions("txlist", "address", addr, from, to, "asc")
}

// Get a list of 'Internal' Transactions by Address
// [Optional Parameters] startblock: starting blockNo to retrieve results, endblock: ending blockNo to retrieve results
// (Returned 'isError' values: 0=No Error, 1=Got Error)
// (Returns up to a maximum of the last 10000 transactions only)

// http://api.etherscan.io/api?module=account&action=txlistinternal&address=0x2c1ba59d6f58433fb1eaee7d20b26ed83bda51a3&startblock=0&endblock=2702578&sort=asc&apikey=YourApiKeyToken

// (To get paginated results use page=<page number> and offset=<max records to return>)
// https://api.etherscan.io/api?module=account&action=txlistinternal&address=0x2c1ba59d6f58433fb1eaee7d20b26ed83bda51a3&startblock=0&endblock=2702578&page=1&offset=10&sort=asc&apikey=YourApiKeyToken

func (a *API) InternalTransactionsByAddress(addr string) (tr TxListRec) {
	// call := "http://api.etherscan.io/api?module=account&action=txlistinternal&address=" + addr + "&startblock=0&endblock=2702578&sort=asc&apikey=" + a.apiKey
	// fmt.Println(call)
	// resp, err := http.Get(call)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return TxListRec{Status: "NOTOK", Message: err.Error()}
	// }
	// err = json.NewDecoder(resp.Body).Decode(&tr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	//http.Error(w, err.Error(), 400)
	// 	return TxListRec{Status: "NOTOK", Message: err.Error()}
	// }
	return a.AccountTransactions("txlistinternal", "address", addr, "0", "2702578", "asc")
}

// Really need to improve the way we call this.
func (a *API) AccountTransactions(action string, by string, val string, start string, end string, sort string) (tr TxListRec) {
	call := "http://api.etherscan.io/api?module=account&action=" +
		action + "&" + by + "=" + val + "&startblock=" + start +
		"&endblock=" + end + "&sort=" + sort + "&apikey=" + a.apiKey
	//fmt.Println(call)
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
