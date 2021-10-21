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
		"\n-10.0\n-9.5\n-9.0\n",
		"\n-0.5\n0.0\n0.5\n",
		"\n9.0\n9.5\n10.0\n",
	}
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	}
}
