package main

import (
	"strings"
	"testing"
)

func TestPrintMonths(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"Month 1: January",
				"Month 2: February",
				"Month 3: March",
				"Month 4: April",
				"Month 5: May",
				"Month 6: June",
				"Month 7: July",
				"Month 8: August",
				"Month 9: September",
				"Month 10: October",
				"Month 11: November",
				"Month 12: December",
				"Month 43: error",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintMonths(writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
