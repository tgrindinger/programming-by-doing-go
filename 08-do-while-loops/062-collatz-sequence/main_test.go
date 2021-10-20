package main

import (
	"strings"
	"testing"
)

func TestPrintCollatz(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "6", "6\n", []string {
				"6\t3\t10\t5\t16\t8\t4\t2\t1",
				"Terminated after 8 steps. The largest value was 16.",
			},
		},
		{ "11", "11\n", []string {
				"11\t34\t17\t52\t26\t13\t40\t20\t10\t5\t16\t8\t4\t2\t1",
				"Terminated after 14 steps. The largest value was 52.",
			},
		},
		{ "27", "27\n", []string {
				"35\t106\t53\t160\t80\t40\t20\t10\t5\t16\t8\t4\t2\t1",
				"Terminated after 111 steps. The largest value was 9232.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintCollatz(reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
