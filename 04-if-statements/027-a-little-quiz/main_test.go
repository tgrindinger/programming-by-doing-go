package main

import (
	"strings"
	"testing"
)

func TestGiveQuiz(t *testing.T) {
	reader := strings.NewReader("N\n3\n1\n2\n")
	writer := &strings.Builder{}
	GiveQuiz(reader, writer)
	got := writer.String()
	want := `Are you ready for a quiz?  Okay, here it comes!

Q1) What is the capital of Alaska?
        1) Melbourne
        2) Anchorage
        3) Juneau

> 
That's right!

Q2) Can you store the value "cat" in a variable of type int?
        1) yes
        2) no

> 
Sorry, "cat" is a string. ints can only store numbers.

Q3) What is the result of 9+6/3?
        1) 5
        2) 11
        3) 15/3

> 
That's right!


Overall, you got 2 out of 3 correct.
Thanks for playing!
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}