package main

import (
	"strings"
	"testing"
)

func TestPrintCount(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		want  string
	}{
		{ "4 13 3", "4\n13\n3\n", "4 7 10 13 \n" },
		{ "5 20 5", "5\n20\n5\n", "5 10 15 20 \n" },
		{ "2 10 1", "2\n10\n1\n", "2 3 4 5 6 7 8 9 10 \n" },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintCount(reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.want) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.want)
			}
		})
	}
}
