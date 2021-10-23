package main

import (
	"strings"
	"testing"
)

func TestPrintName(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		want  string
		num   int
	}{
		{ "gump", "gump\n", "gump", 10 },
		{ "Mitchell", "Mitchell\n", "Mitchell", 5 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintName(reader, writer)
			gots := strings.Split(writer.String(), "\n")
			if count(gots, tc.want) != tc.num {
				t.Errorf("\ngot  %d\nwant %d", count(gots, tc.want), tc.num)
			}
		})
	}
}

func count(arr []string, target string) int {
	count := 0
	for _, s := range arr {
		if s == target {
			count++
		}
	}
	return count
}
