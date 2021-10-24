package main

import (
	"strings"
	"testing"
)

func TestPrintAddresses(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "default", "Esteban\n12\n79.3\nDave\n10\n91\nMichelle\n11\n98.6\n", []string {
				"The names are: Esteban Dave Michelle",
				"The grades are: 12 10 11",
				"The averages are: 79.3 91.0 98.6",
				"The average for the three students is: 89.633333",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintStudents(reader, writer)
			gots := strings.Split(writer.String(), "\n")
			for _, want := range tc.wants {
				if !contains(gots, want) {
					t.Errorf("\ngot  %s\nwant %s", gots, want)
				}
			}
		})
	}
}

func contains(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}
