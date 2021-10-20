package main

import (
	"strings"
	"testing"
)

func TestCheckRightTriangle(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		want  string
	}{
		{ "6 8 10", "6\n8\n10\n", "These sides *do* make a right triangle.  Yippy-skippy!" },
		{ "6 3 8 10", "6\n3\n8\n10\n", "These sides *do* make a right triangle.  Yippy-skippy!" },
		{ "4 5 5", "4\n5\n5\n", "NO!  These sides do not make a right triangle!" },
		{ "4 3 -9 5 1 5", "4\n5\n5\n", "NO!  These sides do not make a right triangle!" },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			CheckRightTriangle(reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.want) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.want)
			}
		})
	}
}