package main

import (
	"strings"
	"testing"
)

func TestRunCalculator(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "addition", "2 + 3\n0 + 2\n", []string {
				"> 5.000000",
				"> Bye, now.",
			},
		},
		{ "subtraction", "5 - 2\n0 + 2\n", []string {
				"> 3.000000",
				"> Bye, now.",
			},
		},
		{ "multiplication", "4 * 9\n0 + 2\n", []string {
				"> 36.000000",
				"> Bye, now.",
			},
		},
		{ "division", "12 / 2\n0 + 2\n", []string {
				"> 6.000000",
				"> Bye, now.",
			},
		},
		{ "modulus", "5 % 2\n0 + 2\n", []string {
				"> 1.000000",
				"> Bye, now.",
			},
		},
		{ "negation", "- 15\n0 + 2\n", []string {
				"> -15.000000",
				"> Bye, now.",
			},
		},
		{ "factorial", "4 !\n0 + 2\n", []string {
				"> 24.000000",
				"> Bye, now.",
			},
		},
		{ "power", "8 ^ 2\n0 + 2\n", []string {
				"> 64.000000",
				"> Bye, now.",
			},
		},
		{ "sin", "sin 10\n0 + 2\n", []string {
				"> -0.544021",
				"> Bye, now.",
			},
		},
		{ "cos", "cos 10\n0 + 2\n", []string {
				"> -0.839072",
				"> Bye, now.",
			},
		},
		{ "tan", "tan 10\n0 + 2\n", []string {
				"> 0.648361",
				"> Bye, now.",
			},
		},
		{ "sqrt", "sqrt 81\n0 + 2\n", []string {
				"> 9.000000",
				"> Bye, now.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			RunCalculator(reader, writer)
			gots := strings.Split(writer.String(), "\n")
			for _, want := range tc.wants {
				if !contains(gots, want) {
					t.Errorf("\ngot  %q\nwant %q", gots, want)
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
