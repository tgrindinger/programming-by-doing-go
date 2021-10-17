package main

import (
	"strings"
	"testing"
)

func TestAskForNumbers(t *testing.T) {
	reader := strings.NewReader("1.1\n2.2\n5.5\n")
	writer := &strings.Builder{}
	AskForNumbers(reader, writer)
	got := writer.String()
	want := `What is your first number? What is your second number? What is your third number? 
( 1.1 + 2.2 + 5.5 ) / 2 is... 4.4
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}