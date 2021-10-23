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
				"Keychains cost $10.00 each.\n",
				"Shipping will be $7.00.\n",
				"Subtotal before tax: $27.00.\n",
				"Tax will be $2.23.\n",
				"Total cost is $29.23.\n",
				"CHECKOUT\n\n",
				"What is your name? ",
				"Thanks for your order, Biff!\n",
			},
		},
		{ "menu error", "5\n1\n3\n4\nBiff\n", 1, []string {
				"Ye Olde Keychain Shoppe\n\n",
				"Please select a valid menu option!\n",
				"\nYou have 0 keychains. How many to add? ",
				"CHECKOUT\n\n",
				"What is your name? ",
				"You have 3 keychains.\n",
				"Keychains cost $10.00 each.\n",
				"Shipping will be $8.00.\n",
				"Subtotal before tax: $38.00.\n",
				"Tax will be $3.14.\n",
				"Total cost is $41.13.\n",
				"Thanks for your order, Biff!\n",
			},
		},
		{ "add error", "1\n-4\n3\n\n4\nBiff\n", 1, []string {
				"Ye Olde Keychain Shoppe\n\n",
				"\nYou have 0 keychains. How many to add? ",
				"You cannot enter a negative number. Try again.\n",
				"You now have 3 keychains.\n\n",
				"CHECKOUT\n\n",
				"What is your name? ",
				"You have 3 keychains.\n",
				"Keychains cost $10.00 each.\n",
				"Shipping will be $8.00.\n",
				"Subtotal before tax: $38.00.\n",
				"Tax will be $3.14.\n",
				"Total cost is $41.13.\n",
				"Thanks for your order, Biff!\n",
			},
		},
		{ "remove error", "1\n3\n2\n4\n1\n4\nBiff\n", 1, []string {
				"Ye Olde Keychain Shoppe\n\n",
				"\nYou have 0 keychains. How many to add? ",
				"You now have 3 keychains.\n\n",
				"\nYou have 3 keychains. How many to remove? ",
				"You don't have that many keychains. Try again: ",
				"You now have 2 keychains.\n\n",
				"CHECKOUT\n\n",
				"What is your name? ",
				"You have 2 keychains.\n",
				"Keychains cost $10.00 each.\n",
				"Shipping will be $7.00.\n",
				"Subtotal before tax: $27.00.\n",
				"Tax will be $2.23.\n",
				"Total cost is $29.23.\n",
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
