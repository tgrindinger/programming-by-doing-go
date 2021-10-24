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
				"\tMake: Toyota",
				"\tModel: Camry",
				"\tYear: 1985",
				"\tLicense: 622-VRX",
				"\tMake: Chevrolet",
				"\tModel: Chevette",
				"\tYear: 1980",
				"\tLicense: J43-SMB",
				"\tMake: Honda",
				"\tModel: Civic",
				"\tYear: 1993",
				"\tLicense: 883-RS9",
				"\tMake: Ford",
				"\tModel: Mustang",
				"\tYear: 1966",
				"\tLicense: AZUCAR",
				"\tMake: Dodge",
				"\tModel: Neon",
				"\tYear: 1996",
				"\tLicense: G74-LLC",
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
