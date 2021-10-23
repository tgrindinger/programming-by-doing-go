package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintFunctions(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		seed  int64
		wants []string
	}{
		{ "default", "jose\n", 1, []string {
				"Here we go.\n\n",
				"\nSome random numbers, if you please: \n",
				"\nFirst: 2\n",
				"\nSecond: 8\n",
				"\nThey were not the same.\n\n",
				"\nMore counting, but this time not by ones: 2 4 6 8 10 8 6 4 2 \n",
				"\nLet's figure a project grade.\n",
				"\n434521 -> 71\n\n",
				"\nFinally, some easy ones.Please enter your name: Hi, jose\n\n",
				"\nDo you feel lucky, punk?\n",
				"\nyou lose\n",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintFunctions(random, reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
