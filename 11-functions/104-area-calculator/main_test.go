package main

import (
	"strings"
	"testing"
)

func TestRunCalculator(t *testing.T) {
	cases := []struct {
		desc string
		input string
		wants []string
	}{
		{ "default", "1\n5\n6\n2\n10\n4\n3\n9\n4\n4\n5\n", []string {
				"The area is 15.000000.",
				"The area is 40.",
				"The area is 81.",
				"The area is 50.265482",
				"Goodbye.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			RunCalculator(reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
