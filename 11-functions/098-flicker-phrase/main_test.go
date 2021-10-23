package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintPhrase(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"I                               \r",
				"  pledge                        \r",
				"         allegiance             \r",
				"                    to the      \r",
				"                           flag.\r",
				"I pledge allegiance to the flag.\n",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			writer := &strings.Builder{}
			PrintPhrase(random, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
