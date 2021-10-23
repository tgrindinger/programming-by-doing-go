package main

import (
	"strings"
	"testing"
)

func TestSellKeychains(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		seed  int64
		wants []string
	}{
		{ "add 3 remove 1", "1\n3\n2\n1\n3\n4\nBiff\n", 1, []string {
				"Ye Olde Keychain Shoppe\n\n",
				"\nYou have 0 keychains. How many to add? ",
				"You now have 3 keychains.\n\n",
				"\nYou have 3 keychains. How many to remove? ",
				"You now have 2 keychains.\n\n",
				"You have 2 keychains.\n",
				"Keychains cost $10 each.\n",
				"Total cost is $20.\n\n",
				"CHECKOUT\n\n",
				"What is your name? ",
				"Thanks for your order, Biff!\n",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			SellKeychains(reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
