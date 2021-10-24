package main

import (
	"strings"
	"testing"
)

func TestPrintPeople(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "nameage", "nameage.txt\n", []string {
				"Steve_Jobs is 50",
				"James_Brown is 71",
				"Britney_Spears is 23",
				"Michael_Jackson is 46",
				"Stanley_Kubrick is 76",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintPeople(reader, writer)
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
