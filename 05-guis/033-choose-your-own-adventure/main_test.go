package main

import (
	"strings"
	"testing"
)

func TestPlayGame(t *testing.T) {
	cases := []struct {
		input    string
		lastLine string
	}{
		{ "kitchen\nrefrigerator\nno\n", "You die of starvation... eventually." },
		{ "kitchen\nrefrigerator\nyes\n", "You die of food poisoning." },
		{ "kitchen\ncabinet\nno\n", "You die of starvation... eventually." },
		{ "kitchen\ncabinet\nyes\n", "You die of food poisoning." },
		{ "upstairs\nbedroom\nno\n", "Well, then I guess you'll never know what was in there.  Thanks for playing, I'm tired of making nexted if statements." },
		{ "upstairs\nbedroom\nyes\n", "A creepy clown jumps out and bonks you on the head." },
		{ "upstairs\nbathroom\nno\n", "Wallow in your filth, I guess." },
		{ "upstairs\nbathroom\nyes\n", "You fail to notice the toaster precariously perched next to the tub until it's too late." },
	}
	for _, tc := range cases {
		t.Run(tc.lastLine, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PlayGame(reader, writer)
			lines := strings.Split(writer.String(), "\n")
			got := lines[len(lines)-2]
			want := tc.lastLine
			if got != want {
				t.Errorf("\ngot  %q\nwant %q", got, want)
			}
		})
	}
}
