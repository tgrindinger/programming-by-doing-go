package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	cases := []struct {
		input  string
		result int
	}{
		{ "6\n9\n-3\n2\n0\n", 14 },
		{ "1\n2\n3\n4\n5\n0\n", 15 },
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d", tc.result), func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			SumNumbers(reader, writer)
			got := writer.String()
			want := fmt.Sprintf("The total is %d.", tc.result)
			if !strings.Contains(got, want) {
				t.Errorf("\ngot  %q\nwant %q", got, want)
			}
		})
	}
}
