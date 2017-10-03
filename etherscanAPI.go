package etherscanAPI

// API - the Etherscan API toolkit
type API struct {
	apiKey string
}

// NewEtherscanAPI - create an api object with the correct key
func NewEtherscanAPI(key string) (newAPI *API) {
	api := API{key}
	newAPI = &api
	return
}
