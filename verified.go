package etherscanAPI

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

func getPageQuery(address string, chain string) (query *browser.Browser, err error) {
	dot := "."
	if len(chain) == 0 {
		dot = ""
	}
	url := "https://" + chain + dot + "etherscan.io/address/" + address + "#code"
	bow := surf.NewBrowser()
	err = bow.Open(url)
	if err == nil {
		query = bow
	}
	return
}

// GetSourceCode returns the code section
func GetSourceCode(address string, chain string) (code string, err error) {
	query, err := getPageQuery(address, chain)
	if err != nil {
		return
	}
	code = query.Find(".js-sourcecopyarea").Text()
	if code[:6] != "pragma" {
		err = errors.New("code not found")
	}
	return
}

// GetABI returns the public ABI
func GetABI(address string, chain string) (abi string, err error) {
	query, err := getPageQuery(address, chain)
	if err != nil {
		return
	}
	abi = query.Find(".js-copytextarea2").Text()
	if abi[:1] != "[" {
		err = errors.New("ABI not found")
	}
	return
}

// GetVersion - return the version
func GetVersion(address string, chain string) (version string, err error) {
	query, err := getPageQuery(address, chain)
	if err != nil {
		return
	}
	err = errors.New("Not Found")
	query.Find("table").Each(func(i int, table *goquery.Selection) {
		next := false
		table.Find("td").Each(func(j int, s *goquery.Selection) {
			if strings.Contains(s.Text(), "Compiler Version:") {
				next = true
			} else if next {
				version = s.Text()
				err = nil
				return
			}
		})
	})
	return
}

// GetOptimisation - find out if the code was optimised
func GetOptimisation(address string, chain string) (optimisation bool, err error) {
	query, err := getPageQuery(address, chain)
	if err != nil {
		return
	}
	err = errors.New("Not Found")
	query.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("td").Each(func(j int, s *goquery.Selection) {
			if strings.Contains(s.Text(), "Optimization") {
				optimisation = strings.Contains(s.Text()[12:], "Enabled")
				err = nil
				return
			}
		})
	})
	return
}

// GetSWARM returns the swarm address
func GetSWARM(address string, chain string) (swarm string, err error) {
	query, err := getPageQuery(address, chain)
	if err != nil {
		return
	}
	err = errors.New("Not Found")
	query.Find("pre").Each(func(i int, pre *goquery.Selection) {
		text := pre.Text()[:7]
		if text == "bzzr://" {
			swarm = pre.Text()
			err = nil
			return
		}
	})
	return
}
