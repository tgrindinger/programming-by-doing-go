package main

import (
	"math/rand"
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
		{ "found", "25\n", []string {
				"ArrayList: [25 15 4 7 16 47 18 28 39 29]",
				"25 is in the ArrayList.",
			}, []string {
				"25 is not in the ArrayList.",
			},
		},
		{ "not found", "26\n", []string {
				"ArrayList: [25 15 4 7 16 47 18 28 39 29]",
			}, []string {
				"26 is in the ArrayList.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			FindElement(random, reader, writer)
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
