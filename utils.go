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
	"fmt"
	"math/big"
	"strings"
)

func decimalStr(bb string, places int) string {
	lBB := len(bb)
	if lBB < places {
		zeros := "000000000000000000"
		bb = "0." + zeros[lBB:] + bb

	} else if lBB == 18 {
		bb = "0." + bb
	} else {
		bb = fmt.Sprintf("%s.%s", bb[:lBB-places], bb[lBB-places:])
	}
	return bb
}

func oneEther() *big.Int {
	return new(big.Int).SetUint64(1000000000000000000)
}

func bi(n int) *big.Int {
	return new(big.Int).SetUint64(uint64(n))
}

func strToWei(value string) (vInWei *big.Int, ok bool) {
	vInWei, ok = new(big.Int).SetString(value, 10)
	return
}

// EtherStrToWei converts 1.0 to 1000000000000000000000
//                        0.1 to  100000000000000000000
func EtherStrToWei(value string) (vInEther *big.Int, ok bool) {
	v, ok := new(big.Int).SetString(value, 10)

	if ok {
		vInEther = new(big.Int).Mul(v, oneEther())
		return
	}
	strA := strings.Split(value, ".")
	if len(strA) != 2 {
		ok = false
		return
	}
	v, ok = new(big.Int).SetString(strA[0], 10)
	if !ok {
		return
	}
	vInWholeEther := new(big.Int).Mul(v, oneEther())
	v2, ok := new(big.Int).SetString(strA[1], 10)
	if !ok {
		return
	}
	pwr := new(big.Int).Exp(bi(10), bi(18-len(strA[1])), nil)
	vInPartEther := new(big.Int).Mul(v2, pwr)
	vInEther = new(big.Int).Add(vInWholeEther, vInPartEther)
	//fmt.Println(strA[0], strA[1])
	return
}
