package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintElements(t *testing.T) {
	cases := []struct {
		desc  string
		wants []string
	}{
		{ "default", []string {
				"64  64  23  56  45  26  17  27  28  58  88  57  39  58  52  16  21  30  20  64  ",
				"15  58  36  70  21  46  32  10  71  62  67  18  15  44  18  83  30  50  69  64  ",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			writer := &strings.Builder{}
			PrintElements(random, writer)
			gots := strings.Split(writer.String(), "\n")
			for _, want := range tc.wants {
				if !contains(gots, want) {
					t.Errorf("\ngot  %s\nwant %s", gots, want)
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
