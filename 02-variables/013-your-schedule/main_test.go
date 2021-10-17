package main

import (
	"strings"
	"testing"
)

func TestPrintSchedule(t *testing.T) {
	var writer strings.Builder
	PrintSchedule(&writer)
	got := writer.String()
	want := `+-------------------------------------------------------------+
| 1 |                           English III |       Ms. Lapan |
| 2 |                           Precalculus |     Mrs. Gideon |
| 3 |                          Music Theory |       Mr. Davis |
| 4 |                         Biotechnology |      Ms. Palmer |
| 5 |            Principles of Technology I |      Ms. Garcia |
| 6 |                              Latin II |    Mrs. Barnett |
| 7 |                         AP US History | Ms. Johannessen |
| 8 | Business Computer Information Systems |       Mr. James |
+-------------------------------------------------------------+
`
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}