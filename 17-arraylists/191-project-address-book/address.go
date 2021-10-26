package main

import (
	"fmt"
	"strings"
)

type Address struct {
	firstname string
	lastname  string
	phone     string
	address   string
	email     string
}

func (a *Address) CompareTo(addy *Address, field string) int {
	result := 0
	switch field {
	case "firstname": result = strings.Compare(a.firstname, addy.firstname)
	case "lastname": result = strings.Compare(a.lastname, addy.lastname)
	case "phone": result = strings.Compare(a.phone, addy.phone)
	case "address": result = strings.Compare(a.address, addy.address)
	case "email": result = strings.Compare(a.email, addy.email)
	}
	return result
}

func (a *Address) Matches(field, value string) bool {
	found := false
	switch field {
	case "firstname": found = strings.Contains(a.firstname, value)
	case "lastname": found = strings.Contains(a.lastname, value)
	case "phone": found = strings.Contains(a.phone, value)
	case "address": found = strings.Contains(a.address, value)
	case "email": found = strings.Contains(a.email, value)
	}
	return found
}

func (a *Address) String() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s", a.firstname, a.lastname, a.phone, a.address, a.email)
}
