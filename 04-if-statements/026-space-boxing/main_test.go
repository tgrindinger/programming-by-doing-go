package main

import (
	"strings"
	"testing"
)

func TestProcessInput(t *testing.T) {
	reader := strings.NewReader("128\n2\n")
	writer := &strings.Builder{}
	ProcessInput(reader, writer)
	got := writer.String()
	want := `Please enter your current earth weight: 
I have information for the following planets:
   1. Venus   2. Mars    3. Jupiter
   4. Saturn  5. Uranus  6. Neptune

Which planet are you visiting? 
Your weight would be 49.92 pounds on that planet.
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
