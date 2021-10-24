package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestSaveCars(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "nameage", "Toyota\nCamry\n1985\n622-VRX\nChevrolet\nChevette\n1980\nJ43-SMB\nHonda\nCivic\n1993\n883-RS9\nFord\nMustang\n1966\nAZUCAR\nDodge\nNeon\n1996\nG74-LLC\ncars.txt\n", []string {
				"Toyota Camry 1985 622-VRX",
				"Chevrolet Chevette 1980 J43-SMB",
				"Honda Civic 1993 883-RS9",
				"Ford Mustang 1966 AZUCAR",
				"Dodge Neon 1996 G74-LLC",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			SaveCars(reader, writer)
			file, _ := os.Open("cars.txt")
			scanner := bufio.NewScanner(file)
			gots := make([]string, 5)
			for i := 0; i < len(gots); i++ {
				scanner.Scan()
				gots[i] = scanner.Text()
			}
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
