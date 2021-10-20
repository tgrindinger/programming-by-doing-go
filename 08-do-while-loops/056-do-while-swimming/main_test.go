package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSwim(t *testing.T) {
	cases := []struct {
		temp  float64
		wants []string
	}{
		{ 80.5, []string { "GALLANT stops swimming. Total swim time: 4 min.", "GOOFUS stops swimming. Total swim time: 4 min." } },
		{ 78, []string { "GALLANT stops swimming. Total swim time: 0 min.", "GOOFUS stops swimming. Total swim time: 1 min." } },
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%.1f", tc.temp), func(t *testing.T) {
			reader := strings.NewReader(fmt.Sprintf("%f\n", tc.temp))
			writer := &strings.Builder{}
			Swim(reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant: %q", got, want)
				}
			}
		})
	}
}
