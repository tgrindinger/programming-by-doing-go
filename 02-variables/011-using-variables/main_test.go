package main

import (
	"strings"
	"testing"
)

func TestPrintVariables(t *testing.T) {
	var writer strings.Builder
	PrintVariables(&writer)
	got := writer.String()
	want := `This is room # 113
e is close to 2.71828
I am learning a bit about Computer Science
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}