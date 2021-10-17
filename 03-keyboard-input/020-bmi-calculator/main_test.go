package main

import (
	"strings"
	"testing"
)

func TestAskForMeasurements(t *testing.T) {
	reader := strings.NewReader("1.75\n73\n")
	writer := &strings.Builder{}
	AskForMeasurements(reader, writer)
	got := writer.String()
	want := `Your height in m: Your weight in kg: 
Your BMI is 23.83673
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}