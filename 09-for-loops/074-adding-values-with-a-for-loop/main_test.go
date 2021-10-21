package main

import (
	"strings"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	cases := []struct {
		desc    string
		input   string
		numbers string
		sum     string
	}{
		{ "5", "5\n", "\n1 2 3 4 5 \n", "The sum is 15." },
		{ "8", "8\n", "\n1 2 3 4 5 6 7 8 \n", "The sum is 36." },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			SumNumbers(reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.numbers) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.numbers)
			}
			if !strings.Contains(got, tc.sum) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.sum)
			}
		})
	}
}
