package main

import (
	"strings"
	"testing"
)

func TestPrintCount(t *testing.T) {
	writer := &strings.Builder{}
	PrintEvenNumbers(writer)
	got := writer.String()
	want := "1\n2 <\n3\n4 <\n5\n6 <\n7\n8 <\n9\n10 <\n11\n12 <\n13\n14 <\n15\n16 <\n17\n18 <\n19\n20 <\n"
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
