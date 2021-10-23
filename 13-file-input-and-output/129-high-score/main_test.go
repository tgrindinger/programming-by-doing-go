package main

import (
	"strings"
	"testing"
)

func TestSaveHighScore(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "default", "32767\nMitchell\n", []string {
				"Mitchell 32767",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			fileWriter := &strings.Builder{}
			SaveHighScore(reader, writer, fileWriter)
			gots := strings.Split(fileWriter.String(), "\n")
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
