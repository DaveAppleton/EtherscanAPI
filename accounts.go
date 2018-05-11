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
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
)

// Get Ether Balance for a single Address
// https://api.etherscan.io/api?module=account&action=balance&address=0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae&tag=latest&apikey=YourApiKeyToken
type balRec struct {
	Status  string
	Message string
	Result  string
}

func (a *API) getEtherBalance(addr string) (res *big.Int, err error) {
	var tr balRec
	var ok bool
	call := "http://api.etherscan.io/api?module=account&action=balance&address=" + addr + "&tag=latest&apikey=" + a.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		return
	}
	if strings.Compare(tr.Status, "1") != 0 {
		err = errors.New(tr.Message)
		return
	}
	res, ok = strToWei(tr.Result)
	if !ok {
		err = errors.New("error in number " + tr.Result)
	}
	return
}

// Get Ether Balance for multiple Addresses in a single call
// https://api.etherscan.io/api?module=account&action=balancemulti&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a,0x63a9975ba31b0b9626b34300f7f627147df1f526,0x198ef1ec325a96cc354c7266a038be8b5c558f67&tag=latest&apikey=YourApiKeyToken

func (a *API) getMultiEtherBalances(addr []string) (tr balRec, err error) {
	//var tr balRec
	addresses := strings.Join(addr, ",")
	call := "http://api.etherscan.io/api?module=account&action=balancemulti&address=" + addresses + "&tag=latest&apikey=" + a.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(tr)
	return
}

type minedBlocks struct {
	fillThisIn string
}

// Get list of Blocks Mined by Address
// https://api.etherscan.io/api?module=account&action=getminedblocks&address=0x9dd134d14d1e65f84b706d6f205cd5b1cd03a46b&blocktype=blocks&apikey=YourApiKeyToken

// func (a *API) getBlocksMinedByAddress(addr string, blocktype string) (tr minedBlocks) {
// 	a.getBlocksMinedByAddressByPage(addr, blocktype, 0, 0)
// }

// (To get paginated results use page=<page number> and offset=<max records to return>)
// ** type = blocks (full blocks only) or uncles (uncle blocks only)
// https://api.etherscan.io/api?module=account&action=getminedblocks&address=0x9dd134d14d1e65f84b706d6f205cd5b1cd03a46b&blocktype=blocks&page=1&offset=10&apikey=YourApiKeyToken

func (a *API) getBlocksMinedByAddressByPage(addr string, blocktype string, page int, itemsPerPage int) (tr minedBlocks) {
	//var tr minedBlocks
	pageData := ""
	if itemsPerPage != 0 {
		pageData = fmt.Sprintf("page=%d&offset=%d", page, itemsPerPage)
	}
	call := "https://api.etherscan.io/api?module=account&action=getminedblocks&address=" + addr + "&blocktype=" + blocktype + "&apikey=" + a.apiKey + pageData
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tr)
	return
}
