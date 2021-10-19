package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintMessage(t *testing.T) {
	cases := []struct {
		seed    int64
		message string
	}{
		{ 0, "MAGIC 8-BALL SAYS: Concentrate and ask again\n" },
		{ 1, "MAGIC 8-BALL SAYS: It is decidedly so\n" },
		{ 2, "MAGIC 8-BALL SAYS: Most likely\n" },
		{ 3, "MAGIC 8-BALL SAYS: Signs point to yes\n" },
	}
	for _, tc := range cases {
		t.Run(tc.message, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			writer := &strings.Builder{}
			PrintMessage(random, writer)
			got := writer.String()
			if got != tc.message {
				t.Errorf("\ngot  %q\nwant %q", got, tc.message)
			}
		})
	}
}