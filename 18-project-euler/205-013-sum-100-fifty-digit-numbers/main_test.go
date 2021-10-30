package main

import (
	"math/big"
	"testing"
)

func TestSumBigInts(t *testing.T) {
	bigInts := readBigInts(bigStrings)
	got := SumBigInts(bigInts)
	want, _ := new(big.Int).SetString("5537376230390876637302048746832985971773659831892672", 0)
	if got.Cmp(want) != 0 {
		t.Errorf("\ngot  %s\nwant %s", got, want)
	}
}
