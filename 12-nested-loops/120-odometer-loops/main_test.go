package main

import (
	"strings"
	"testing"
)

func TestPrintCounter(t *testing.T) {
	cases := []struct {
		desc     string
		input    string
		ms       int64
		wants    []string
		notWants []string
	}{
		{ "default", "10\n", 0, []string {
				" 0001",
				" 0010",
				" 0100",
				" 1000",
				" 0009",
				" 0099",
				" 0999",
				" 9999",
			}, []string {
			},
		},
		{ "default", "8\n", 0, []string {
				" 0001",
				" 0010",
				" 0100",
				" 1000",
				" 0007",
				" 0077",
				" 0777",
				" 7777",
			}, []string {
				" 0009",
				" 0099",
				" 0999",
				" 9999",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintCounter(reader, writer, tc.ms)
			gots := strings.Split(writer.String(), "\r")
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
