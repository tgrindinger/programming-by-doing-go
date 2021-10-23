package main

import (
	"strings"
	"testing"
)

func TestPrintTable(t *testing.T) {
	cases := []struct {
		desc     string
		wants    []string
	}{
		{ "default", []string {
				"1 | 1\t2\t3\t4\t5\t6\t7\t8\t9\t",
				"2 | 2\t4\t6\t8\t10\t12\t14\t16\t18\t",
				"3 | 3\t6\t9\t12\t15\t18\t21\t24\t27\t",
				"4 | 4\t8\t12\t16\t20\t24\t28\t32\t36\t",
				"5 | 5\t10\t15\t20\t25\t30\t35\t40\t45\t",
				"6 | 6\t12\t18\t24\t30\t36\t42\t48\t54\t",
				"7 | 7\t14\t21\t28\t35\t42\t49\t56\t63\t",
				"8 | 8\t16\t24\t32\t40\t48\t56\t64\t72\t",
				"9 | 9\t18\t27\t36\t45\t54\t63\t72\t81\t",
				"10 | 10\t20\t30\t40\t50\t60\t70\t80\t90\t",
				"11 | 11\t22\t33\t44\t55\t66\t77\t88\t99\t",
				"12 | 12\t24\t36\t48\t60\t72\t84\t96\t108\t",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintTable(writer)
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
