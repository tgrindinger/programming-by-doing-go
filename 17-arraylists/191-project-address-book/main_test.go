package main

import (
	"io"
	"strings"
	"testing"
)

type StubFileOpener struct {
	reader io.Reader
	writer io.Writer
}

func (f *StubFileOpener) OpenReader(filename string) io.Reader {
	return f.reader
}

func (f *StubFileOpener) OpenWriter(filename string) io.Writer {
	return f.writer
}

func (f *StubFileOpener) CloseReader() { }

func (f *StubFileOpener) CloseWriter() { }

func TestManipulateAddressBook(t *testing.T) {
	cases := []struct {
		desc      string
		input     string
		fileInput string
		wants     []string
		notWants  []string
		fileWants []string
	}{
		{ "add, save", "3\nsome\nguy\n123-456-7890\n123 nowhere blvd, smalltown, ks 12345\nsomeguy@domain.com\n2\naddsave.txt\n8\n", "", []string {
			}, []string {
			}, []string {
				"some|guy|123-456-7890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com",
			},
		},
		{ "add, remove, save", "3\nsome\nguy\n123-456-7890\n123 nowhere blvd, smalltown, ks 12345\nsomeguy@domain.com\n4\n1\n2\naddsave.txt\n8\n", "", []string {
			}, []string {
			}, []string {
				"",
			},
		},
		{ "add, edit, save", "3\nsome\nguy\n123-456-7890\n123 nowhere blvd, smalltown, ks 12345\nsomeguy@domain.com\n5\n1\nthe\nguy\n123-456-7890\n123 nowhere blvd, smalltown, ks 12345\nsomeguy@domain.com\n2\naddsave.txt\n8\n", "", []string {
			}, []string {
			}, []string {
				"the|guy|123-456-7890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com",
			},
		},
		{ "load, save", "1\ndummy.txt\n2\ndummy.txt\n8\n", 
				"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com\nother|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org", []string {
			}, []string {
			}, []string {
				"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com",
				"other|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org",
			},
		},
		{ "load, sort, save", "1\ndummy.txt\n6\n1\n2\ndummy.txt\n8\n", 
				"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com\nother|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org", []string {
			}, []string {
			}, []string {
				"other|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org",
				"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com",
			},
		},
		{ "load, search", "1\ndummy.txt\n7\n1\nsome\n8\n", 
				"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com\nother|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org",
			[]string {
				"Result: some, guy, 1234567890, 123 nowhere blvd, smalltown, ks 12345, someguy@domain.com",
			}, []string {
			}, []string {
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			fileReader := strings.NewReader(tc.fileInput)
			fileWriter := &strings.Builder{}
			addybook := NewAddressBook(&StubFileOpener{fileReader, fileWriter})
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			manipulateAddressBook(reader, writer, addybook)
			gots := strings.Split(writer.String(), "\n")
			assertGot(t, "stdout", gots, tc.wants)
			assertNotGot(t, "stdout", gots, tc.notWants)
			fileGots := strings.Split(fileWriter.String(), "\n")
			assertGot(t, "file", fileGots, tc.fileWants)
		})
	}
}

func assertNotGot(t testing.TB, label string, gots []string, notWants []string) {
	t.Helper()
	for i := 0; i < len(notWants); i++ {
		if contains(gots, notWants[i]) {
			t.Errorf("\n%s\ngot     %v\nnotWant %s", label, gots, notWants[i])
		}
	}
}

func assertGot(t testing.TB, label string, gots []string, wants []string) {
	t.Helper()
	current := 0
	for i := 0; i < len(gots); i++ {
		if current == len(wants) {
			return
		}
		if gots[i] == wants[current] {
			current++
		}
	}
	if current < len(wants) {
		t.Errorf("\n%s\ngot  %v\nwant %v", label, gots, wants)
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
