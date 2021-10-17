package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("making an important message", func(t *testing.T) {
		got := Hello()
		want := "Mr. Mitchell is cool."
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
