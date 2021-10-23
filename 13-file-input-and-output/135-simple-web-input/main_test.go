package main

import (
	"strings"
	"testing"
)

func TestPrintWebLine(t *testing.T) {
	cases := []struct {
		desc  string
		wants []string
	}{
		{ "SimpleWebInput.java", []string {
				"import java.net.URL;",
				"import java.util.Scanner;",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintWebLine(writer)
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
