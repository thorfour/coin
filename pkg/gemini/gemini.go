package gemini

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	api    = "https://api.gemini.com/v1/pubticker/"
	btcusd = "btcusd"
	ethusd = "ethusd"
)

// Get returns a BTCUSD & ETHUSD quote
func Get() (map[string]*Quote, error) {

	var err error
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// Get BTCUSD quote
	var btcquote *Quote
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		btcquote, err = GetSymbol(btcusd)
	}(wg)
	// Get ETHUSD quote
	var ethquote *Quote
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		ethquote, err = GetSymbol(ethusd)
	}(wg)

	wg.Wait()

	return map[string]*Quote{
		btcusd: btcquote,
		ethusd: ethquote,
	}, err
}

// GetSymbol returns a quote for the given symbol
func GetSymbol(symbol string) (*Quote, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", api, symbol))
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
