package main

import "testing"

func TestInitials(t *testing.T) {
	got := Initials()
	want := `TTTTT JJJJJ  GGG
  T     J   G   G
  T     J   G
  T     J   GGGGG
  T   J J   G   G
  T   J J   G   G
  T    JJ    GGG
`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
