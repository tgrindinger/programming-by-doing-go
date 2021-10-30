package main

import (
	"math/big"
	"testing"
)

// pass
// 10 01
// fail
// 00 11

// pass
// 1100 1010 1001 0110 0101 0011
// fail
// 1000 0100 0010 0001
// 1110 1101 1011 0111
// 1111

// pass
// 111000 110100 110010 110001
// 101100 101010 101001
// 100110 100101 100011
// 011100 011010 011001
// 010110 010101 010011
// 001110 001101 001011
// 000111
// fail
// 100000 010000 001000 000100 000010 000001
// 110000 101000 100100 100010 100001
//

func TestNumberPaths(t *testing.T) {
	cases := []struct{
		desc  string
		input int
		want  *big.Int
	}{
		{ "1x1", 1, big.NewInt(2) },
		{ "2x2", 2, big.NewInt(6) },
		{ "3x3", 3, big.NewInt(20) },
		{ "20x20", 20, big.NewInt(137846528820) },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := numberPaths(tc.input)
			if got.Cmp(tc.want) != 0 {
				t.Errorf("\ngot  %s\nwant %s", got, tc.want)
			}
		})
	}
}
