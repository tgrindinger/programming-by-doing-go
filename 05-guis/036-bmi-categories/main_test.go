package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	cases := []struct {
		height   float64
		weight   float64
		bmi      string
		category string
	}{
		{ 69, 220, "32.5", "moderately obese" },
		{ 70, 90, "12.9", "very severely underweight" },
	}
	for _, tc := range cases {
		t.Run(tc.category, func(t *testing.T) {
			reader := strings.NewReader(fmt.Sprintf("%f\n%f\n", tc.height, tc.weight))
			writer := &strings.Builder{}
			AskQuestions(reader, writer)
			lines := strings.Split(writer.String(), "\n")
			bmiLine := lines[len(lines)-3]
			categoryLine := lines[len(lines)-2]
			if !strings.Contains(bmiLine, tc.bmi) {
				t.Errorf("got %q want %q", bmiLine, tc.bmi)
			}
			if !strings.Contains(categoryLine, tc.category) {
				t.Errorf("got %q want %q", categoryLine, tc.category)
			}
		})
	}
}
