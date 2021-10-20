package main

import (
	"strings"
	"testing"
)

func TestDoAdventure(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		want  string
	}{
		{ "eat food", "kitchen\nback\nupstairs\ndownstairs\nkitchen\nrefrigerator\nyes\n", "You have died." },
		{ "take a nap", "kitchen\nback\nupstairs\nbedroom\nnap\n", "You have a pleasant nap." },
		{ "toasted", "kitchen\nback\nupstairs\nbathroom\nbath\n", "You are crispy." },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			DoAdventure(reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.want) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.want)
			}
		})
	}
}
