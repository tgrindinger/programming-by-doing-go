package main

import (
	"strings"
	"testing"
)

func TestPrintCount(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "short", "short", []string {
				"Your message is 5 characters long.",
				"The first character is at position 0 and is 's'.",
				"The last character is at position 4 and is 't'.",
				"Your message contains the letter 'a' 0 times. Isn't that interesting?",
			},
		},
		{ "panama", "A man, a plan, a canal: Panama!\n", []string {
				"Your message is 31 characters long.",
				"The first character is at position 0 and is 'A'.",
				"The last character is at position 30 and is '!'.",
				"Your message contains the letter 'a' 10 times. Isn't that interesting?",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintMessage(reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
