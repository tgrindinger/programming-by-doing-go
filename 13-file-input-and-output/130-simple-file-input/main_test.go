package main

import (
	"strings"
	"testing"
)

func TestPrintName(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "default", "Horace Mann\n", []string {
				"Using my psychic powers (aided by reading data from the file), I have determined that your name is Horace Mann.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintName(reader, writer)
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
