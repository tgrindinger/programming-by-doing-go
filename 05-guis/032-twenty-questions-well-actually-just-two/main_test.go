package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	cases := []struct {
		input string
		guess string
	}{
		{ "animal\nno\n", "squirrel" },
		{ "animal\nyes\n", "moose" },
		{ "vegetable\nno\n", "carrot" },
		{ "vegetable\nyes\n", "watermelon" },
		{ "mineral\nno\n", "paper clip" },
		{ "mineral\nyes\n", "Camaro" },
	}
	for _, tc := range cases {
		t.Run(tc.guess, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			AskQuestions(reader, writer)
			got := writer.String()
			want := `TWO QUESTIONS!
Think of an object, and I'll try to guess it.

Question 1) Is it animal, vegetable, or mineral?
> 
Question 2) Is it bigger than a breadbox?
> 
My guess is that you are thinking of a %s.
I would ask you if I'm right, but I don't actually care.
`
			want = fmt.Sprintf(want, tc.guess)
			if got != want {
				t.Errorf("\ngot  %q\nwant %q", got, want)
			}
		})
	}
}
