package main

import (
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	t.Run("ask dennis", func(t *testing.T) {
		reader := strings.NewReader("Dennis\n37\n8.50\n")
		writer := &strings.Builder{}
		AskQuestions(reader, writer)
		got := writer.String()
		want := `Hello.  What is your name?

Hi, Dennis!  How old are you?

So you're 37, eh?  That's not old at all!
How much do you make, Dennis?

8.5!  I hope that's per hour and not per year! LOL!
`
		if got != want {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	})

	t.Run("ask catsup", func(t *testing.T) {
		reader := strings.NewReader("Catsup\n12\n99.9\n")
		writer := &strings.Builder{}
		AskQuestions(reader, writer)
		got := writer.String()
		want := `Hello.  What is your name?

Hi, Catsup!  How old are you?

So you're 12, eh?  That's not old at all!
How much do you make, Catsup?

99.9!  I hope that's per hour and not per year! LOL!
`
		if got != want {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	})
}
