package main

import (
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	reader := strings.NewReader("Billy_Corgan\n17\n")
	writer := &strings.Builder{}
	AskQuestions(reader, writer)
	got := writer.String()
	want := `Hey, what's your name? 
Ok, Billy_Corgan, how old are you? 
You can't vote, Billy_Corgan.
You can't rent a car, Billy_Corgan.
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
