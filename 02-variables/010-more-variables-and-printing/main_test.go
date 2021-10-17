package main

import (
	"strings"
	"testing"
)

func TestPrintVariables(t *testing.T) {
	var writer strings.Builder
	PrintVariables(&writer)
	got := writer.String()
	want := `Let's talk about Zed A. Shaw.
He's 74 inches tall.
He's 180 pounds heavy.
Actually that's not too heavy.
He's got Blue eyes and Brown hair.
His teeth are usually White depending on the coffee.
If I add 35, 74, and 180 I get 289.
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}