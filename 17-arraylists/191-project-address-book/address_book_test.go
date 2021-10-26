package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSave(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc     string
		addys    []*Address
		wants    []string
	}{
		{ "first", []*Address{addy1}, []string {
				"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com",
			},
		},
		{ "second", []*Address{addy2}, []string {
				"other|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org",
			},
		},
		{ "both", []*Address{addy1, addy2}, []string {
				"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com",
				"other|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			addybook := &AddressBook{tc.addys, &StubFileOpener{nil, writer}}
			addybook.Save("dummy.txt")
			gots := strings.Split(writer.String(), "\n")
			assertGot(t, tc.desc, gots, tc.wants)
		})
	}
}

func TestLoad(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc  string
		addys []*Address
		input string
		want  string
	}{
		{ "first", []*Address{addy1},
			"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com",
			"[some, guy, 1234567890, 123 nowhere blvd, smalltown, ks 12345, someguy@domain.com]",
		},
		{ "second", []*Address{addy2},
			"other|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org",
			"[other, dude, 3216540987, 456 somewhere ave, bigcity, ca 54321, otherdude@email.org]",
		},
		{ "both", []*Address{addy1, addy2},
			"some|guy|1234567890|123 nowhere blvd, smalltown, ks 12345|someguy@domain.com\nother|dude|3216540987|456 somewhere ave, bigcity, ca 54321|otherdude@email.org",
			"[some, guy, 1234567890, 123 nowhere blvd, smalltown, ks 12345, someguy@domain.com other, dude, 3216540987, 456 somewhere ave, bigcity, ca 54321, otherdude@email.org]",
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			addybook := &AddressBook{tc.addys, &StubFileOpener{reader, nil}}
			addybook.Load("dummy.txt")
			got := addybook.String()
			if got != tc.want {
				t.Errorf("\ngot  %s\nwant %s", got, tc.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc     string
		addybook *AddressBook
		index    int
		want     int
	}{
		{ "one", &AddressBook{[]*Address{addy1}, nil}, 0, 0 },
		{ "two", &AddressBook{[]*Address{addy1, addy2}, nil}, 1, 1 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.addybook.Remove(tc.index)
			got := len(tc.addybook.addresses)
			if got != tc.want {
				t.Errorf("got %d want %d", got, tc.want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc     string
		addybook *AddressBook
		addy     *Address
		index    int
		want     int
	}{
		{ "addy1 -> addy2", &AddressBook{[]*Address{addy1}, nil}, addy2, 0, 1 },
		{ "addy2 -> addy1", &AddressBook{[]*Address{addy2}, nil}, addy1, 0, 1 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.addybook.Replace(tc.index, tc.addy)
			got := len(tc.addybook.addresses)
			if got != tc.want {
				t.Errorf("got %d want %d", got, tc.want)
			}
			if tc.addybook.addresses[0] != tc.addy {
				t.Errorf("got %v want %v", tc.addybook.addresses[0], tc.addy)
			}
		})
	}
}

func TestSort(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc     string
		addybook *AddressBook
		field    string
		want     *AddressBook
	}{
		{ "firstname", &AddressBook{[]*Address{addy1, addy2}, nil}, "firstname", &AddressBook{[]*Address{addy2, addy1}, nil} },
		{ "firstname reversed", &AddressBook{[]*Address{addy2, addy1}, nil}, "firstname", &AddressBook{[]*Address{addy2, addy1}, nil} },
		{ "lastname", &AddressBook{[]*Address{addy1, addy2}, nil}, "lastname", &AddressBook{[]*Address{addy2, addy1}, nil} },
		{ "lastname reversed", &AddressBook{[]*Address{addy2, addy1}, nil}, "lastname", &AddressBook{[]*Address{addy2, addy1}, nil} },
		{ "phone", &AddressBook{[]*Address{addy1, addy2}, nil}, "phone", &AddressBook{[]*Address{addy1, addy2}, nil} },
		{ "phone reversed", &AddressBook{[]*Address{addy2, addy1}, nil}, "phone", &AddressBook{[]*Address{addy1, addy2}, nil} },
		{ "address", &AddressBook{[]*Address{addy1, addy2}, nil}, "address", &AddressBook{[]*Address{addy1, addy2}, nil} },
		{ "address reversed", &AddressBook{[]*Address{addy2, addy1}, nil}, "address", &AddressBook{[]*Address{addy1, addy2}, nil} },
		{ "email", &AddressBook{[]*Address{addy1, addy2}, nil}, "email", &AddressBook{[]*Address{addy2, addy1}, nil} },
		{ "email reversed", &AddressBook{[]*Address{addy2, addy1}, nil}, "email", &AddressBook{[]*Address{addy2, addy1}, nil} },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.addybook.Sort(tc.field)
			for i, addy := range tc.addybook.addresses {
				if addy != tc.want.addresses[i] {
					t.Errorf("got %v want %v", addy, tc.want.addresses[i])
				}
			}
		})
	}
}

func TestSearch(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc     string
		addybook *AddressBook
		field    string
		value    string
		want     *Address
	}{
		{ "firstname addy1", &AddressBook{[]*Address{addy1, addy2}, nil}, "firstname", "some", addy1 },
		{ "firstname addy2", &AddressBook{[]*Address{addy1, addy2}, nil}, "firstname", "other", addy2 },
		{ "lastname addy1", &AddressBook{[]*Address{addy1, addy2}, nil}, "lastname", "guy", addy1 },
		{ "lastname addy2", &AddressBook{[]*Address{addy1, addy2}, nil}, "lastname", "dude", addy2 },
		{ "phone addy1", &AddressBook{[]*Address{addy1, addy2}, nil}, "phone", "123", addy1 },
		{ "phone addy2", &AddressBook{[]*Address{addy1, addy2}, nil}, "phone", "321", addy2 },
		{ "address addy1", &AddressBook{[]*Address{addy1, addy2}, nil}, "address", "123", addy1 },
		{ "address addy2", &AddressBook{[]*Address{addy1, addy2}, nil}, "address", "456", addy2 },
		{ "email addy1", &AddressBook{[]*Address{addy1, addy2}, nil}, "email", "some", addy1 },
		{ "email addy2", &AddressBook{[]*Address{addy1, addy2}, nil}, "email", "other", addy2 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			addy := tc.addybook.Search(tc.field, tc.value)
			if addy != tc.want {
				t.Errorf("got %v want %v", addy, tc.want)
			}
		})
	}
}

func TestPretty(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc     string
		addybook *AddressBook
		want     string
	}{
		{ "empty", &AddressBook{}, "" },
		{ "one", &AddressBook{[]*Address{addy1}, nil}, fmt.Sprintf("1) %v\n", addy1) },
		{ "two", &AddressBook{[]*Address{addy1, addy2}, nil}, fmt.Sprintf("1) %v\n2) %v\n", addy1, addy2) },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.addybook.Pretty()
			if got != tc.want {
				t.Errorf("\ngot  %s\nwant %s", got, tc.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc     string
		addybook *AddressBook
		addy     *Address
		want     int
	}{
		{ "empty", &AddressBook{}, addy1, 1 },
		{ "one", &AddressBook{[]*Address{addy1}, nil}, addy2, 2 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.addybook.Add(tc.addy)
			got := len(tc.addybook.addresses)
			if got != tc.want {
				t.Errorf("got %d want %d", got, tc.want)
			}
		})
	}
}

func TestAddressBookString(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc     string
		addybook *AddressBook
		want     string
	}{
		{ "empty", &AddressBook{}, "[]" },
		{ "one", &AddressBook{[]*Address{addy1}, nil}, fmt.Sprintf("[%v]", addy1) },
		{ "two", &AddressBook{[]*Address{addy1, addy2}, nil}, fmt.Sprintf("[%v %v]", addy1, addy2) },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.addybook.String()
			if got != tc.want {
				t.Errorf("\ngot  %s\nwant %s", got, tc.want)
			}
		})
	}
}