package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrintSquareRoot(t *testing.T) {
	cases := []struct {
		num   int
		input string
		want  string
	}{
		{ 9, "9\n", "The square root of 9 is 3.000000." },
		{ 2, "2\n", "The square root of 2 is 1.414214." },
		{ 10, "-9\n-10\n10\n", "The square root of 10 is 3.162278." },
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d", tc.num), func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintSquareRoot(reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.want) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.want)
			}
		})
	}
}
