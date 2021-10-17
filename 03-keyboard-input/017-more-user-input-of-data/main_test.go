package main

import (
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	reader := strings.NewReader("Helena\nBonham-Carter\n12\n453916\nbonham_453916\n3.73\n")
	writer := &strings.Builder{}
	AskQuestions(reader, writer)
	got := writer.String()
	want := `Please enter the following information so I can sell it for a profit!

First name: Last name: Grade (9-12): Student ID: Login: GPA (0.0-4.0): 
Your information:
        Login: bonham_453916
        ID:    453916
        Name:  Bonham-Carter, Helena
        GPA:   3.73
        Grade: 12
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
