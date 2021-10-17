package main

import (
	"strings"
	"testing"
)

func TestCommentedLines(t *testing.T) {
	var writer strings.Builder
	PrintCommentedLines(&writer)
	got := writer.String()
	want := `I could have code like this.
This will run.
`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
