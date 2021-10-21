package main

import (
	"strings"
	"testing"
)

func TestPrintCount(t *testing.T) {
	writer := &strings.Builder{}
	PrintCount(writer)
	got := writer.String()
	want := "15\n18\n21\n24\n27\n30\n"
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
