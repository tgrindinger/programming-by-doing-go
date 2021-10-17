package main

import (
	"strings"
	"testing"
)

func TestPrintNameAndGraduationYear(t *testing.T) {
	var writer strings.Builder
	PrintNameAndGraduationYear(&writer)
	got := writer.String()
	want := "My name is Juan Valdez and I'll graduate in 2010.\n"
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}