package main

import (
	"strings"
	"testing"
)

func TestPlayGame(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "fight", "1\n", []string {
				"Misty sent out STARMIE!",
				"What will PIKACHU do?",
				"PIKACHU used THUNDERSHOCK!  It's super effective.",
				"STARMIE used WATER PULSE!",
				"PIKACHU fainted.",
			},
		},
		{ "swap", "2\n", []string {
				"Misty sent out STARMIE!",
				"What will PIKACHU do?",
				"PIKACHU swaps out with GYARADOS!",
				"GYARADOS used BITE!  It's super effective.",
				"STARMIE fainted.",
			},
		},
		{ "run", "3\n", []string {
				"Misty sent out STARMIE!",
				"What will PIKACHU do?",
				"PIKACHU couldn't run. Skip a turn, coward.",
				"STARMIE used WATER PULSE!",
				"PIKACHU fainted.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PlayGame(reader, writer)
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
