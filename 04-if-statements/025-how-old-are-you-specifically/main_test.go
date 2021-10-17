package main

import (
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	t.Run("billy", func(t *testing.T) {
		reader := strings.NewReader("Billy_Corgan\n17\n")
		writer := &strings.Builder{}
		AskQuestions(reader, writer)
		got := writer.String()
		want := `Hey, what's your name? (Sorry, I keep forgetting.) Ok, Billy_Corgan, how old are you? 
You can drive but not vote, Billy_Corgan.
`
		if got != want {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	})

	t.Run("caspar", func(t *testing.T) {
		reader := strings.NewReader("Caspar\n14\n")
		writer := &strings.Builder{}
		AskQuestions(reader, writer)
		got := writer.String()
		want := `Hey, what's your name? (Sorry, I keep forgetting.) Ok, Caspar, how old are you? 
You can't drive, Caspar.
`
		if got != want {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	})
}
