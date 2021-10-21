package main

import (
	"strings"
	"testing"
)

func TestPrintFizzBuzz(t *testing.T) {
	writer := &strings.Builder{}
	PrintFizzBuzz(writer)
	got := writer.String()
	wants := []string {
		"1\n2\nFizz\n4\nBuzz\nFizz\n",
		"13\n14\nFizzBuzz\n16\n17\nFizz\n",
		"97\n98\nFizz\nBuzz\n",
	}
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	}
}
