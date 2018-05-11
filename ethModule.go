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
	"math/big"
	"net/http"
)

// https://api.etherscan.io/api?module=proxy&action=eth_getCode&address=0xf75e354c5edc8efed9b59ee9f67a80845ade7d0c&tag=latest&apikey=YourApiKeyToken

// GetCode - get code at a particular address
func (a *API) GetCode(address string) (res StringRec) {
	// GasPrice - get optimal gas price
	var tr StringRec
	call := "https://api.etherscan.io/api?module=proxy&action=eth_getCode&address=" + address + "&tag=latest&apikey=" + a.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return StringRec{Status: "NOTOK", Message: err.Error()}
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), 400)
		return StringRec{Status: "NOTOK", Message: err.Error()}
	}
	return
}

// IntRec - integer or nothing
type IntRec struct {
	Status  string
	Message string
	Result  *big.Int
}

// StringRec - integer or nothing
type StringRec struct {
	Status  string
	Message string
	Result  string
}

// https://api.etherscan.io/api?module=proxy&action=eth_gasPrice&apikey=YourApiKeyToken

// GasPrice - get optimal gas price
func (a *API) GasPrice() (res IntRec) {
	var tr IntRec
	call := "https://api.etherscan.io/api?module=proxy&action=eth_gasPrice&apikey=&apikey=" + a.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return IntRec{Status: "NOTOK", Message: err.Error()}
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), 400)
		return IntRec{Status: "NOTOK", Message: err.Error()}
	}
	return

}

// https://api.etherscan.io/api?module=proxy&action=eth_estimateGas&to=0xf0160428a8552ac9bb7e050d90eeade4ddd52843&value=0xff22&gasPrice=0x051da038cc&gas=0xffffff&apikey=YourApiKeyToken

// EstimateGas - run Tx locally and return amount of gas required
func (a *API) EstimateGas(to string, value string) (res IntRec) {
	var tr IntRec
	call := "https://api.etherscan.io/api?module=proxy&action=eth_estimateGas&to=" + to + "&value=" + value + "&gasPrice=0x051da038cc&gas=0xffffff&apikey=" + a.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return IntRec{Status: "NOTOK", Message: err.Error()}
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), 400)
		return IntRec{Status: "NOTOK", Message: err.Error()}
	}
	return

}

func (a *API) EthCall(to string, data string) (res StringRec) {
	var sr StringRec
	call := "http://api.etherscan.io/api?module=proxy&action=eth_call&to="
	call += to
	call += "&data="
	call += data
	call += "&tag=latest&apikey="
	call += a.apiKey
	//fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return StringRec{Status: "NOTOK", Message: err.Error()}
	}
	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), 400)
		return StringRec{Status: "NOTOK", Message: err.Error()}
	}
	return sr
}
