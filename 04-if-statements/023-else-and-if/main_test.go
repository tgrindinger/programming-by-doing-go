package main

import (
	"strings"
	"testing"
)

func TestPrintIfs(t *testing.T) {
	writer := &strings.Builder{}
	PrintIfs(writer)
	got := writer.String()
	want := `We should take the cars.
Maybe we could take the buses.
All right, let's just take the buses.
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
