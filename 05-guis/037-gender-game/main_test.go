package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	cases := []struct {
		gender    string
		firstname string
		lastname  string
		age       int
		married   string
		result    string
	}{
		{ "F", "Kim", "Kardashian", 32, "y\n", "Mrs. Kardashian." },
		{ "F", "Marisa", "Tomei", 48, "n\n", "Ms. Tomei." },
		{ "F", "Chloe", "Moretz", 16, "", "Chloe Moretz" },
		{ "M", "Daniel", "Radcliffe", 23, "", "Mr. Radcliffe" },
		{ "M", "Zachary", "Gordon", 15, "", "Zachary Gordon" },
	}
	for _, tc := range cases {
		t.Run(tc.result, func(t *testing.T) {
			reader := strings.NewReader(fmt.Sprintf("%s\n%s\n%s\n%d\n%s", tc.gender, tc.firstname, tc.lastname, tc.age, tc.married))
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