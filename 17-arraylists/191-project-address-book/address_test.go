package main

import (
	"testing"
)

func Addy1() *Address {
	return &Address{"some", "guy", "1234567890", "123 nowhere blvd, smalltown, ks 12345", "someguy@domain.com"}
}

func Addy2() *Address {
	return &Address{"other", "dude", "3216540987", "456 somewhere ave, bigcity, ca 54321", "otherdude@email.org"}
}

func TestCompareTo(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc  string
		field string
		a1    *Address
		a2    *Address
		want  int
	}{
		{ "firstname", "firstname", addy1, addy2, 1 },
		{ "firstname reverse", "firstname", addy2, addy1, -1 },
		{ "lastname", "lastname", addy1, addy2, 1 },
		{ "lastname reverse", "lastname", addy2, addy1, -1 },
		{ "phone", "phone", addy1, addy2, -1 },
		{ "phone reverse", "phone", addy2, addy1, 1 },
		{ "address", "address", addy1, addy2, -1 },
		{ "address reverse", "address", addy2, addy1, 1 },
		{ "email", "email", addy1, addy2, 1 },
		{ "email reverse", "email", addy2, addy1, -1 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.a1.CompareTo(tc.a2, tc.field)
			if got != tc.want {
				t.Errorf("got %d want %d", got, tc.want)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc  string
		field string
		value string
		want  bool
	}{
		{ "firstname true", "firstname", addy1.firstname, true },
		{ "firstname false", "firstname", addy2.firstname, false },
		{ "lastname true", "lastname", addy1.lastname, true },
		{ "lastname false", "lastname", addy2.lastname, false },
		{ "phone true", "phone", addy1.phone, true },
		{ "phone false", "phone", addy2.phone, false },
		{ "address true", "address", addy1.address, true },
		{ "address false", "address", addy2.address, false },
		{ "email true", "email", addy1.email, true },
		{ "email false", "email", addy2.email, false },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := addy1.Matches(tc.field, tc.value)
			if got != tc.want {
				t.Errorf("got %t want %t", got, tc.want)
			}
		})
	}
}

func TestAddressString(t *testing.T) {
	addy1 := Addy1()
	addy2 := Addy2()
	cases := []struct {
		desc  string
		addy  *Address
		want  string
	}{
		{ "addy1", addy1, "some, guy, 1234567890, 123 nowhere blvd, smalltown, ks 12345, someguy@domain.com" },
		{ "addy2", addy2, "other, dude, 3216540987, 456 somewhere ave, bigcity, ca 54321, otherdude@email.org" },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.addy.String()
			if got != tc.want {
				t.Errorf("\ngot  %s\nwant %s", got, tc.want)
			}
		})
	}
}
