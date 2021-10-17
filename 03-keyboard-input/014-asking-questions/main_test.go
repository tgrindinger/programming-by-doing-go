package main

import (
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	scanner := strings.NewReader("35\n6'2\"\n180\n")
	writer := &strings.Builder{}
	AskQuestions(scanner, writer)
	got := writer.String()
	want := `How old are you? How tall are you? How much do you weigh? So you're 35 old, 6'2" tall and 180 heavy.`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}