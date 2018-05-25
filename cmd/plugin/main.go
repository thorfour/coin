package main

import (
	"net/url"

	"github.com/thorfour/coin/pkg/gemini"
)

// Handler implements the sillyputty plugin model
func Handler(_ url.Values) (string, error) {
	q, err := gemini.Get()
	if err != nil {
		return "", err
	}

	return q.Last, nil
}
