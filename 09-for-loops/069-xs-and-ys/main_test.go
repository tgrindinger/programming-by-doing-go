package main

import (
	"strings"
	"testing"
)

func TestPrintCount(t *testing.T) {
	writer := &strings.Builder{}
	PrintCount(writer)
	got := writer.String()
	wants := []string {
		"\n-10.0\t100.00\n-9.5\t90.25\n-9.0\t81.00\n", 
		"\n-0.5\t0.25\n0.0\t0.00\n0.5\t0.25\n", 
		"\n9.0\t81.00\n9.5\t90.25\n10.0\t100.00\n", 
	}
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	}
}
