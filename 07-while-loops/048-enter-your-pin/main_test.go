package main

import (
	"strings"
	"testing"
)

func TestGetPin(t *testing.T) {
	reader := strings.NewReader("90210\n11111\n12345\n")
	writer := &strings.Builder{}
	GetPin(reader, writer)
	got := writer.String()
	want1 := "INCORRECT PIN. TRY AGAIN."
	want2 := "PIN ACCEPTED. YOU NOW HAVE ACCESS TO YOUR ACCOUNT."
	if !strings.Contains(got, want1) {
		t.Errorf("\ngot  %q\nwant %q", got, want1)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("\ngot  %q\nwant %q", got, want2)
	}
}
