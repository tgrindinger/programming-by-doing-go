package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintArrayList(t *testing.T) {
	cases := []struct {
		desc     string
		wants    []string
		notWants []string
	}{
		{ "default", []string {
				"ArrayList before: [64 64 23 56 45 26 17 27 28 58 88 57 39 58 52 16 21 30 20 64]",
				"ArrayList after : [16 17 20 21 23 26 27 28 30 39 45 52 56 57 58 58 64 64 64 88]",
			}, []string {
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			writer := &strings.Builder{}
			FindElement(random, writer)
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
