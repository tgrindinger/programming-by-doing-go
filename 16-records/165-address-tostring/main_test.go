package main

import (
	"strings"
	"testing"
)

func TestPrintAddresses(t *testing.T) {
	cases := []struct {
		desc  string
		wants []string
	}{
		{ "default", []string {
				"1. 191 Marigold Lane, Miami, FL 33179",
				"2. 3029 Losh Lane, Crafton, PA 15205",
				"3. 4939 Holt Street, Lake Worth, FL 33463",
				"4. 2693 Hannah Street, Hickory, NC 28601",
				"5. 4880 Carter Street, Fairview Heights, IL 62208",
				"6. 2682 Austin Avenue, Douglas, GA 31533",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintAddresses(writer)
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
