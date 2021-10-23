package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintFunctions(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"Watch as we demonstrate functions.\n\n",
				"\nI'm going to get a random character from A-Z\n",
				"\nThe character is: C\n\n",
				"\nNow let's count from -10 to 10\n",
				"\n-10 -9 -8 -7 -6 -5 -4 -3 -2 -1 0 1 2 3 4 5 6 7 8 9 10 How was that?\n",
				"\nNow we take the absolute value of a number.\n",
				"\n|-10| = 10\n\n",
				"That's all.  This program has been brought to you by:\n\n",
				"\nprogrammed by Graham Mitchell\n",
				"\nmodified by Thomas Grindinger\n",
				"\nThis code is distributed under the terms of the standard BSD license.  Do with it as you wish.\n",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			writer := &strings.Builder{}
			PrintFunctions(random, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
