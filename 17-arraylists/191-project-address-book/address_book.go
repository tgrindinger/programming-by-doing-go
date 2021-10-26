package main

import (
	"bufio"
	"fmt"
	"strings"
)

type AddressBook struct {
	addresses  []*Address
	fileOpener IFileOpener
}

func NewAddressBook(fileOpener IFileOpener) *AddressBook {
	return &AddressBook{[]*Address{}, fileOpener}
}

func (a *AddressBook) Load(filename string) {
	a.addresses = []*Address{}
	file := a.fileOpener.OpenReader(filename)
	defer a.fileOpener.CloseReader()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "|")
		a.addresses = append(a.addresses, &Address{values[0], values[1], values[2], values[3], values[4]})
	}
}

func (a *AddressBook) Save(filename string) {
	file := a.fileOpener.OpenWriter(filename)
	defer a.fileOpener.CloseWriter()
	for _, addy := range a.addresses {
		fmt.Fprintf(file, "%s\n", fmt.Sprintf("%s|%s|%s|%s|%s", addy.firstname, addy.lastname, addy.phone, addy.address, addy.email))
	}
}

func (a *AddressBook) Add(address *Address) {
	a.addresses = append(a.addresses, address)
}

func (a *AddressBook) Remove(entry int) {
	if entry == len(a.addresses) - 1 {
		a.addresses = a.addresses[:entry]
	} else {
		a.addresses = append(a.addresses[:entry], a.addresses[entry+1:]...)
	}
}

func (a *AddressBook) Replace(entry int, address *Address) {
	a.addresses[entry] = address
}

func (a *AddressBook) Sort(sortBy string) {
	for i := 0; i < len(a.addresses); i++ {
		for j := i + 1; j < len(a.addresses); j++ {
			if a.addresses[i].CompareTo(a.addresses[j], sortBy) > 0 {
				temp := a.addresses[i]
				a.addresses[i] = a.addresses[j]
				a.addresses[j] = temp
			}
		}
	}
}

func (a *AddressBook) Search(field, value string) *Address {
	for _, addy := range a.addresses {
		if addy.Matches(field, value) {
			return addy
		}
	}
	return nil
}

func (a *AddressBook) String() string {
	return fmt.Sprint(a.addresses)
}

func (a *AddressBook) Pretty() string {
	builder := strings.Builder{}
	for i, addy := range a.addresses {
		builder.WriteString(fmt.Sprintf("%d) %s\n", i+1, addy))
	}
	return builder.String()
}
