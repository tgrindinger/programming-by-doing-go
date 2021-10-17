package main

import (
	"strings"
	"testing"
)

func TestPrintVariables(t *testing.T) {
	var writer strings.Builder
	PrintVariables(&writer)
	got := writer.String()
	want :=
		`There are 100 cars available.
There are only 30 drivers available.
There will be 70 empty cars today.
We can transport 120.0 people today.
We have 90 to carpool today.
We need to put about 3.0 in each car.
`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
