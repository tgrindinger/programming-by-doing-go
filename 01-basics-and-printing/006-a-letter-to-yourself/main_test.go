package main

import "testing"

func TestLetter(t *testing.T) {
	got := Letter()
	want := `+-------------------------------------------------------+
|                                                  #### |
|                                                  #### |
|                                                  #### |
|                                                       |
|                                                       |
|                        Bill Gates                     |
|                        1 Microsoft Way                |
|                        Redmond, WA 98104              |
|                                                       |
+-------------------------------------------------------+
`
	if got != want {
		t.Errorf("got %q want %v", got, want)
	}
}
