package main

import (
	"fmt"
	"os"

	"github.com/thorfour/coin/pkg/gemini"
)

func main() {
	q, err := gemini.Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to obtain quote")
		return
	}

	var resp string
	for s, t := range q {
		resp = fmt.Sprintf("%s%s: %v\n", resp, s, t.Last)
	}

	fmt.Print(resp)
}
