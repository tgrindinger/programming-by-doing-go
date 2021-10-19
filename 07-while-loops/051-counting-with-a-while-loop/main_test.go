package main

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	reader := strings.NewReader("All work and no play makes Jack a dull boy.\n3\n")
	writer := &strings.Builder{}
	Count(reader, writer)
	got := writer.String()
	wants := []string {
		"10. All work and no play makes Jack a dull boy.",
		"20. All work and no play makes Jack a dull boy.",
		"30. All work and no play makes Jack a dull boy.",
	}
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	}
}
