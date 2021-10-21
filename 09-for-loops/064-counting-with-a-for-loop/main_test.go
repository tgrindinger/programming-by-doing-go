package main

import (
	"strings"
	"testing"
)

func TestPrintMessage(t *testing.T) {
	reader := strings.NewReader("Hey, hey.\n")
	writer := &strings.Builder{}
	PrintMessage(reader, writer)
	got := writer.String()
	want := "2. Hey, hey.\n4. Hey, hey.\n6. Hey, hey.\n8. Hey, hey.\n10. Hey, hey.\n12. Hey, hey.\n14. Hey, hey.\n16. Hey, hey.\n18. Hey, hey.\n20. Hey, hey.\n"
	if !strings.Contains(got, want) {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
