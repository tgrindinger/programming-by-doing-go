package main

import (
	"strings"
	"testing"
)

func TestPrintAreas(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"A triangle with sides 3,3,3 has an area of 3.897114",
				"A triangle with sides 3,4,5 has an area of 6.000000",
				"A triangle with sides 7,8,9 has an area of 26.832816",
				"A triangle with sides 5,12,13 has an area of 30.000000",
				"A triangle with sides 10,9,11 has an area of 42.426407",
				"A triangle with sides 8,15,17 has an area of 60.000000",
				"A triangle with sides 9,9,9 has an area of 35.074029",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintAreas(writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
