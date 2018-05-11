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

// https://api.etherscan.io/api?module=logs&action=getLogs
// &fromBlock=379224
// &toBlock=latest
// &address=0x33990122638b9132ca29c723bdf037f1a891a70c
// &topic0=0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545
// &apikey=YourApiKeyToken

// https://api.etherscan.io/api?module=logs&action=getLogs
// &fromBlock=379224
// &toBlock=400000
// &address=0x33990122638b9132ca29c723bdf037f1a891a70c
// &topic0=0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545
// &topic0_1_opr=and
// &topic1=0x72657075746174696f6e00000000000000000000000000000000000000000000
// &apikey=YourApiKeyToken

// EvListItem - returned by an EventList
type EvListItem struct {
	Address          string
	Topics           []string
	Data             string
	BlockNumber      string
	TimeStamp        string
	GasPrice         string
	GasUsed          string
	LogIndex         string
	TransactionHash  string
	TransactionIndex string
}

// EventRec -
type EventRec struct {
	Status  string
	Message string
	Result  []EvListItem
}

// GetLogs - scour event logs from start to stop
func (a *API) GetLogs(fromBlock string, toBlock string, address string, topics []string, operands []string) (ev EventRec) {
	topicLen := len(topics)
	opLen := len(operands)
	if topicLen < 1 {
		ev = EventRec{Status: "NOTOK", Message: "No topics supplied"}
		return
	}
	if topicLen > 3 {
		ev = EventRec{Status: "NOTOK", Message: "Too many topics supplied (3 max)"}
		return
	}
	if topicLen-opLen != 1 {
		ev = EventRec{Status: "NOTOK", Message: "Not enough operands for topics supplied"}
		return
	}

	baseURL := fmt.Sprintf("http://api.etherscan.io/api?module=logs&action=getLogs&fromBlock=%s&toBlock=%s&address=%s&apikey=%s",
		fromBlock,
		toBlock,
		address,
		a.apiKey)

	for i := 0; i < topicLen; i++ {
		if i > 0 {
			baseURL += fmt.Sprintf("&topic%d_%d_opr=%s", i-1, i, operands[i-1])
		}
		baseURL += fmt.Sprintf("&topic%d=%s", i, topics[i])
	}
	fmt.Println(baseURL)
	resp, err := http.Get(baseURL)
	if err != nil {
		return EventRec{Status: "NOTOK", Message: err.Error()}
	}
	err = json.NewDecoder(resp.Body).Decode(&ev)
	if err != nil {
		ev = EventRec{Status: "NOTOK", Message: err.Error()}
	}
	return
}
