package main

import (
	"strings"
	"testing"
)

func TestLookupId(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "Zimmerman", "307760\n", []string {
				"\tMitchell 99.5 123456",
				"\tOrtiz 78.5 813225",
				"\tLuu 95.6 823669",
				"\tZimmerman 96.8 307760",
				"\tBrooks 82.7 827131",
				"Found in slot 3",
				"\tName: Zimmerman",
				"\tAverage: 96.8",
				"\tID: 307760",
			},
		},
		{ "Mitchell", "123456\n", []string {
				"Found in slot 0",
				"\tName: Mitchell",
				"\tAverage: 99.5",
				"\tID: 123456",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			LookupId(reader, writer)
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
