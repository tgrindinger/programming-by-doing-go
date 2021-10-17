package main

import (
	"strings"
	"testing"
)

func TestPrintMath(t *testing.T) {
	var builder strings.Builder
	PrintMath(&builder)
	got := builder.String()
	want := `I will now count my chickens:
Hens 30
Roosters 97
Now I will count the eggs:
7
Is it true that 3 + 2 < 5 - 7?
false
What is 3 + 2? 5
What is 5 - 7? -2
Oh, that's why it's false.
How about some more.
Is it greater? true
Is it greater or equal? true
Is it less or equal? false
`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
