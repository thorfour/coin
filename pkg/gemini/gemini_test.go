package gemini

import (
	"fmt"
	"testing"
)

func TestQuote(t *testing.T) {
	q, err := Get()
	if err != nil {
		t.Fatalf("Err: %v", err)
	}

	fmt.Println(q)
}
