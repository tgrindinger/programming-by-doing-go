package main

import (
	"strings"
	"testing"
)

func TestPrintSum(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "3nums", "3nums.txt\n", []string {
				"3 + 1 + 4 = 8",
			},
		},
		{ "3nums2", "3nums2.txt\n", []string {
				"6 + 3 + 3 = 12",
			},
		},
		{ "3nums4", "3nums4.txt\n", []string {
				"7 + 2 + 5 = 14",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintSum(reader, writer)
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
