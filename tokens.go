package etherscanAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
)

func (a *API) tokenSupply(addr string) (val *big.Int, err error) {
	var tr balRec
	call := "http://api.etherscan.io/api?module=stats&action=tokensupply&contractaddress=" + addr + "&apikey=" + a.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), 400)
		return
	}
	if strings.Compare(tr.Status, "1") != 0 {
		err = errors.New(tr.Message)
		return
	}
	var ok bool
	val, ok = new(big.Int).SetString(tr.Result, 10)
	if !ok {
		err = errors.New("error understanding " + tr.Result)
	}
	return
}

func (a *API) tokenAccountBalance(addr string, account string) (tr balRec, err error) {
	call := "http://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=" + addr + "&address=" + account + "&tag=latest&apikey=" + a.apiKey
	//fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), 400)
		return
	}
	//fmt.Println(tr)
	return
}
