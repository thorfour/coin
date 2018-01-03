package gemini

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	api    = "https://api.gemini.com/v1/pubticker/"
	btcusd = "btcusd"
	ethusd = "ethusd"
)

// Get returns a BTCUSD quote
func Get() (*Quote, error) {

	// Get BTCUSD quote
	resp, err := http.Get(fmt.Sprintf("%s%s", api, btcusd))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	j, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	q := new(Quote)
	err = json.Unmarshal(j, q)
	if err != nil {
		return nil, err
	}

	return q, nil
}
