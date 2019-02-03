package coin

import (
	"fmt"
	"net/url"

	"github.com/thorfour/coin/pkg/gemini"
)

// SillyPuttyHandler implements a handler for a sillyputty server
func SillyPuttyHandler(_ url.Values) (string, error) {
	q, err := gemini.Get()
	if err != nil {
		return "", err
	}

	var resp string
	for s, i := range q {
		resp = fmt.Sprintf("%s:%s: %v\n", resp, s, i.Last)
	}

	return resp, nil
}
