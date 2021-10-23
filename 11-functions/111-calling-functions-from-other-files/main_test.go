package main

import (
	"strings"
	"testing"
)

func TestPrintWeekday(t *testing.T) {
	cases := []struct {
		desc string
		input string
		wants []string
	}{
		{ "11 11 2010", "11 11 2010", []string {
				"12 10 2003 => Wednesday, December 10, 2003",
				" 2 13 1976 => Friday, February 13, 1976",
				" 2 13 1977 => Sunday, February 13, 1977",
				" 7  2 1974 => Tuesday, July 2, 1974",
				" 1 15 2003 => Wednesday, January 15, 2003",
				"10 13 2000 => Friday, October 13, 2000",
				"You were born on Thursday, November 11, 2010!",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintWeekday(reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
