package main

import (
	"strings"
	"testing"
)

func TestPrintCars(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "cars", "cars.txt\n", []string {
				"ArrayList: [1966 Ford Mustang (AZUCAR) 1980 Chevrolet Chevette (J43-SMB) 1985 Toyota Camry (622-VRX) 1993 Honda Civic (883-RS9) 1996 Dodge Neon (G74-LLC)]",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintCars(reader, writer)
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
