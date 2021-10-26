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
				"ArrayList before: [ask not what your country can do for you but]",
				"ArrayList after : [ask but can country do for not what you your]",
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
