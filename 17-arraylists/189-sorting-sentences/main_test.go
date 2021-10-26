package main

import (
	"strings"
	"testing"
)

func TestPrintArrayList(t *testing.T) {
	cases := []struct {
		desc     string
		input    string
		wants    []string
		notWants []string
	}{
		{ "default", "How I want a drink alcoholic of course after the heavy lectures involving Quantum Mechanics\n", []string {
				"Sentence: Sorted: [a after alcoholic course drink heavy how i involving lectures mechanics of quantum the want]",
			}, []string {
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			FindElement(reader, writer)
			gots := strings.Split(writer.String(), "\n")
			for _, want := range tc.wants {
				if !contains(gots, want) {
					t.Errorf("\ngot  %s\nwant %s", gots, want)
				}
			}
			for _, notWant := range tc.notWants {
				if contains(gots, notWant) {
					t.Errorf("\ngot     %s\nnotWant %s", gots, notWant)
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
