package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	cases := []struct {
		guess  int
		result string
	}{
		{ 3, "W00T!  U SUX0R!!!  I PWN J00!!!  IT WAS 4!" },
		{ 4, "LOL!!! U GOT IT!  I CANT BELEIVE U GESSED IT WAS 4!" },
	}
	for _, tc := range cases {
		t.Run(tc.result, func(t *testing.T) {
			reader := strings.NewReader(fmt.Sprintf("%d\n", tc.guess))
			writer := &strings.Builder{}
			GuessNumber(reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.result) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.result)
			}
		})
	}
}
