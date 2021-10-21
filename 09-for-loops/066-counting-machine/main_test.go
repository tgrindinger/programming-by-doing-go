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
		{ "8", "8\n", "0 1 2 3 4 5 6 7 8 \n" },
		{ "19", "19\n", "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 \n" },
		{ "25", "25\n", "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 \n" },
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