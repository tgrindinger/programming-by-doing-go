package main

import "testing"

func TestPrintStuff(t *testing.T) {
	got := Lines()
	want := `Hello World!
Hello Again
I like typing this.
This is fun.
Yay! Printing.
I'd much rather you 'not'.
I "said" do not touch this.`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
