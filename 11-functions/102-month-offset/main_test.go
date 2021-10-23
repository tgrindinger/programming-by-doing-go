package main

import (
	"strings"
	"testing"
)

func TestPrintMonthOffsets(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"Offset for month 1: 1",
				"Offset for month 2: 4",
				"Offset for month 3: 4",
				"Offset for month 4: 0",
				"Offset for month 5: 2",
				"Offset for month 6: 5",
				"Offset for month 7: 0",
				"Offset for month 8: 3",
				"Offset for month 9: 6",
				"Offset for month 10: 1",
				"Offset for month 11: 4",
				"Offset for month 12: 6",
				"Offset for month 43: -1",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintMonthOffsets(writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
