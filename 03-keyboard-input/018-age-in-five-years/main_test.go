package main

import (
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	t.Run("ask percy", func(t *testing.T) {
		reader := strings.NewReader("Percy_Bysshe_Shelley\n34\n")
		writer := &strings.Builder{}
		AskQuestions(reader, writer)
		got := writer.String()
		want := `Hello.  What is your name? 
Hi, Percy_Bysshe_Shelley!  How old are you? 
Did you know that in five years you will be 39 years old?
And five years ago you were 29! Imagine that!
`
		if got != want {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	})

	t.Run("ask Gramps", func(t *testing.T) {
		reader := strings.NewReader("Gramps\n87\n")
		writer := &strings.Builder{}
		AskQuestions(reader, writer)
		got := writer.String()
		want := `Hello.  What is your name? 
Hi, Gramps!  How old are you? 
Did you know that in five years you will be 92 years old?
And five years ago you were 82! Imagine that!
`
		if got != want {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	})
}
