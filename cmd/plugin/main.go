package main

import (
	"net/url"

	"github.com/thorfour/coin/pkg/coin"
)

// Handler implements the sillyputty plugin model
func Handler(_ url.Values) (string, error) {
	return coin.SillyPuttyHandler(nil)
}
