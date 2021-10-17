package main

import (
	"strings"
	"testing"
)

func TestPrintIfs(t *testing.T) {
	writer := &strings.Builder{}
	PrintIfs(writer)
	got := writer.String()
	want := `Too many cats!  The world is doomed!
The world is dry!
People are greater than or equal to dogs.
People are less than or equal to dogs.
People are dogs.
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}