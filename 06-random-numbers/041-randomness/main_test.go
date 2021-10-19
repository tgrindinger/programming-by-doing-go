package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintRandomNumbers(t *testing.T) {
	random := rand.New(rand.NewSource(0))
	writer := &strings.Builder{}
	PrintRandomNumbers(random, writer)
	got := writer.String()
	want := `My random number is 5
Here are some numbers from 1 to 5!
5 4 2 1 2 3 
Here are some numbers from 1 to 100!
78	89	29	69	48	60	
The random numbers were different! Not too surprising, actually.
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}