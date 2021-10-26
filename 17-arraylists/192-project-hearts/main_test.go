package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPlayGame(t *testing.T) {
	cases := []struct {
		desc  string
		seed  int64
		input string
		wants []string
	}{
		{ "basic", 0, "13\n13\n13\n13\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n9\n8\n4\n5\n3\n3\n8\n4\n7\n7\n7\n7\n6\n6\n3\n6\n5\n5\n2\n5\n4\n2\n4\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n13\n13\n13\n5\n12\n12\n12\n12\n11\n7\n6\n11\n1\n1\n1\n1\n1\n1\n1\n1\n1\n3\n3\n2\n1\n1\n1\n3\n1\n1\n3\n2\n1\n3\n2\n2\n1\n1\n1\n1\n1\n1\n2\n1\n1\n1\n1\n1\n1\n1\n1\n1\n13\n13\n13\n1\n1\n1\n1\n3\n1\n1\n1\n1\n1\n1\n1\n1\n9\n9\n9\n9\n1\n2\n4\n3\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n13\n13\n13\n1\n1\n1\n1\n1\n1\n1\n1\n10\n5\n7\n9\n1\n9\n4\n6\n8\n1\n4\n2\n1\n1\n1\n3\n1\n1\n2\n1\n1\n1\n2\n1\n1\n1\n1\n2\n1\n1\n1\n3\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n", []string {
				"Game over. Player 1 wins!",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			playGame(random, reader, writer)
			gots := strings.Split(writer.String(), "\n")
			assertGot(t, "stdout", gots, tc.wants)
		})
	}
}

func assertGot(t testing.TB, label string, gots []string, wants []string) {
	t.Helper()
	current := 0
	for i := 0; i < len(gots); i++ {
		if current == len(wants) {
			return
		}
		if gots[i] == wants[current] {
			current++
		}
	}
	if current < len(wants) {
		t.Errorf("\n%s\ngot  %v\nwant %v", label, gots, wants)
	}
}
