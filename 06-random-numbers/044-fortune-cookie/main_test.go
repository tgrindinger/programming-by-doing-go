package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestGiveFortune(t *testing.T) {
	cases := []struct {
		seed    int64
		fortune string
	}{
		{0, "You will find happiness with a new love.\n    19 - 14 - 29 - 18 - 17 - 8"},
		{1, "Face the truth with dignity.\n    16 - 30 - 12 - 32 - 7 - 8"},
		{2, "Respect your elders. You could inherit a large sum of money.\n    43 - 49 - 15 - 21 - 27 - 29"},
		{4, "You have good reason to be self-confident.\n    11 - 50 - 24 - 10 - 44 - 41"},
		{7, "Someone thinks the world of you.\n    13 - 46 - 10 - 39 - 15 - 29"},
		{12, "Your smile lights up someone else's day.\n    38 - 23 - 19 - 33 - 25 - 16"},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d - %s", tc.seed, tc.fortune), func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			writer := &strings.Builder{}
			GiveFortune(random, writer)
			got := writer.String()
			if !strings.Contains(got, tc.fortune) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.fortune)
			}
		})
	}
}
