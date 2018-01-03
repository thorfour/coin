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

	fmt.Println(q.Last)
}
