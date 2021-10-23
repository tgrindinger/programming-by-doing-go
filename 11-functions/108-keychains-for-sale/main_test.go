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
		{ "add view", "1\n3\n4\n", 1, []string {
				"Ye Olde Keychain Shoppe\n\n",
				"ADD KEYCHAINS\n\n",
				"VIEW ORDER\n\n",
				"CHECKOUT\n",
			},
		},
		{ "add remove view", "1\n2\n3\n4\n", 1, []string {
				"Ye Olde Keychain Shoppe\n\n",
				"ADD KEYCHAINS\n\n",
				"REMOVE KEYCHAINS\n\n",
				"VIEW ORDER\n\n",
				"CHECKOUT\n",
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
