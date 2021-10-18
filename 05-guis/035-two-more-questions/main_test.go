package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	cases := []struct {
		inside string
		alive  string
		result string
	}{
		{ "inside", "yes", "houseplant" },
		{ "inside", "no", "shower curtain" },
		{ "outside", "yes", "bison" },
		{ "outside", "no", "billboard" },
		{ "both", "yes", "dog" },
		{ "both", "no", "cell phone" },
	}
	for _, tc := range cases {
		t.Run(tc.result, func(t *testing.T) {
			reader := strings.NewReader(fmt.Sprintf("%s\n%s\n", tc.inside, tc.alive))
			writer := &strings.Builder{}
			AskQuestions(reader, writer)
			lines := strings.Split(writer.String(), "\n")
			got := lines[len(lines)-2]
			if !strings.Contains(got, tc.result) {
				t.Errorf("got %q want %q", got, tc.result)
			}
		})
	}
}