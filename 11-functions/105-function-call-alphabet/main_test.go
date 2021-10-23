package main

import (
	"strings"
	"testing"
)

func TestPrintFunctionAlphabet(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"Ant Banana Crocodile Doggie Elephant Frog Gorilla Horseradish ",
				"Ice_cream Jackrabbit Kiwi Lhasa_apso Monkey! Narwhal Orangutan Parrot Quail ",
				"Rabbit Snake Thyme Ugli_fruit Valentine_candy Walrus X_men Yams Zebra ",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintFunctionAlphabet(writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
