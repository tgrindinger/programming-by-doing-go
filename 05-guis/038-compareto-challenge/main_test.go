package main

import (
	"strings"
	"testing"
)

func TestPrintComparisons(t *testing.T) {
	cases := []struct {
		first  string
		second string
		result string
	}{
		{ "axe", "dog", "Comparing \"axe\" with \"dog\" produces -1\n" },
		{ "bug", "buzz", "Comparing \"bug\" with \"buzz\" produces -1\n" },
		{ "cat", "frog", "Comparing \"cat\" with \"frog\" produces -1\n" },
		{ "deer", "giant", "Comparing \"deer\" with \"giant\" produces -1\n" },
		{ "money", "stuff", "Comparing \"money\" with \"stuff\" produces -1\n" },
		{ "beaver", "beaver", "Comparing \"beaver\" with \"beaver\" produces 0\n" },
		{ "custard", "custard", "Comparing \"custard\" with \"custard\" produces 0\n" },
		{ "bee", "artichoke", "Comparing \"bee\" with \"artichoke\" produces 1\n" },
		{ "frank", "betty", "Comparing \"frank\" with \"betty\" produces 1\n" },
		{ "zap", "apple", "Comparing \"zap\" with \"apple\" produces 1\n" },
		{ "tool", "charm", "Comparing \"tool\" with \"charm\" produces 1\n" },
		{ "things", "fun", "Comparing \"things\" with \"fun\" produces 1\n" },
	}
	for _, tc := range cases {
		t.Run(tc.result, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintComparison(writer, tc.first, tc.second)
			got := writer.String()
			if got != tc.result {
				t.Errorf("\ngot  %q\nwant %q", got, tc.result)
			}
		})
	}
}