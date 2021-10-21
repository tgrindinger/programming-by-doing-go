package main

import (
	"strings"
	"testing"
)

func TestPrintMessage(t *testing.T) {
	writer := &strings.Builder{}
	PrintMessage(writer)
	got := writer.String()
	want := "1. Mr. Mitchell is cool.\n2. Mr. Mitchell is cool.\n3. Mr. Mitchell is cool.\n4. Mr. Mitchell is cool.\n5. Mr. Mitchell is cool.\n6. Mr. Mitchell is cool.\n7. Mr. Mitchell is cool.\n8. Mr. Mitchell is cool.\n9. Mr. Mitchell is cool.\n10. Mr. Mitchell is cool.\n"
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}