package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAskQuestion(t *testing.T) {
	cases := []struct {
		name   string
		result string
	}{
		{ "Banter", "You don't have to wait long, \"Banter\"." },
		{ "Carswell", "You don't have to wait long, \"Carswell\"." },
		{ "Hope", "That's not bad, \"Hope\"." },
		{ "Jones", "That's not bad, \"Jones\"." },
		{ "Martin", "Looks like a bit of a wait, \"Martin\"." },
		{ "Smith", "Looks like a bit of a wait, \"Smith\"." },
		{ "Travis", "It's gonna be a while, \"Travis\"." },
		{ "Young", "It's gonna be a while, \"Young\"." },
		{ "Zelda", "Not going anywhere for a while, \"Zelda\"?" },
	}
	for _, tc := range cases {
		t.Run(tc.result, func(t *testing.T) {
			reader := strings.NewReader(fmt.Sprintf("%s\n", tc.name))
			writer := &strings.Builder{}
			AskQuestion(reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.result) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.result)
			}
		})
	}
}
