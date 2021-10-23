package main

import (
	"strings"
	"testing"
)

func TestPrintFile(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "3nums", "3nums.txt\n", []string {
				"Which file would you like to read numbers from: Reading numbers from \"3nums.txt\"",
				"3 1 4 ",
				"Total is 8",
			},
		},
		{ "4nums", "4nums.txt\n", []string {
				"Which file would you like to read numbers from: Reading numbers from \"4nums.txt\"",
				"6 3 3 2 ",
				"Total is 14",
			},
		},
		{ "5nums", "5nums.txt\n", []string {
				"Which file would you like to read numbers from: Reading numbers from \"5nums.txt\"",
				"7 2 5 6 1 ",
				"Total is 21",
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
