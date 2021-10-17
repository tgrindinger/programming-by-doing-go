package main

import (
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	reader := strings.NewReader("lorry!\ndeoxyribonucleic?\n42\n1\n")
	writer := &strings.Builder{}
	AskQuestions(reader, writer)
	got := writer.String()
	want := `Give me a word!
Give me a second word!

Great, now your favorite number?
And your second-favorite number...

Whew!  Wasn't that fun?
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
