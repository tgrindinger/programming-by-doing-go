package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func loadAddressBook(scanner *bufio.Scanner, writer io.Writer, addybook *AddressBook) {
	fmt.Fprint(writer, "Which file contains the address book: ")
	scanner.Scan()
	filename := scanner.Text()
	addybook.Load(filename)
	fmt.Fprintf(writer, "Address book loaded from %s\n", filename)
}

func saveAddressBook(scanner *bufio.Scanner, writer io.Writer, addybook *AddressBook) {
	fmt.Fprint(writer, "Which file should contain the address book: ")
	scanner.Scan()
	filename := scanner.Text()
	addybook.Save(filename)
	fmt.Fprintf(writer, "Address book saved to %s\n", filename)
}

func readEntry(scanner *bufio.Scanner, writer io.Writer, addybook *AddressBook) *Address {
	addy := &Address{}
	fmt.Fprint(writer, "First name: ")
	scanner.Scan()
	addy.firstname = scanner.Text()
	fmt.Fprint(writer, "Last name: ")
	scanner.Scan()
	addy.lastname = scanner.Text()
	fmt.Fprint(writer, "Phone number: ")
	scanner.Scan()
	addy.phone = scanner.Text()
	fmt.Fprint(writer, "Address: ")
	scanner.Scan()
	addy.address = scanner.Text()
	fmt.Fprint(writer, "Email: ")
	scanner.Scan()
	addy.email = scanner.Text()
	return addy
}

func addEntry(scanner *bufio.Scanner, writer io.Writer, addybook *AddressBook) {
	addy := readEntry(scanner, writer, addybook)
	addybook.Add(addy)
	fmt.Fprintln(writer, "Entry added.")
}

func removeEntry(scanner *bufio.Scanner, writer io.Writer, addybook *AddressBook) {
	fmt.Fprint(writer, addybook.Pretty())
	fmt.Fprint(writer, "Entry to remove: ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	addybook.Remove(choice - 1)
	fmt.Fprintln(writer, "Entry removed.")
}

func editEntry(scanner *bufio.Scanner, writer io.Writer, addybook *AddressBook) {
	fmt.Fprint(writer, addybook.Pretty())
	fmt.Fprint(writer, "Entry to edit: ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	addy := readEntry(scanner, writer, addybook)
	addybook.Replace(choice - 1, addy)
	fmt.Fprintln(writer, "Entry changed.")
}

func sort(scanner *bufio.Scanner, writer io.Writer, addybook *AddressBook) {
	fmt.Fprintln(writer, "1) First name")
	fmt.Fprintln(writer, "2) Last name")
	fmt.Fprintln(writer, "3) Phone number")
	fmt.Fprintln(writer, "4) Address")
	fmt.Fprintln(writer, "5) Email")
	fmt.Fprint(writer, "Field to sort on: ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer, "Sorting address book...")
	switch choice {
	case 1: addybook.Sort("firstname")
	case 2: addybook.Sort("lastname")
	case 3: addybook.Sort("phone")
	case 4: addybook.Sort("address")
	case 5: addybook.Sort("email")
	}
	fmt.Fprintln(writer, "Finished sorting address book.")
}

func search(scanner *bufio.Scanner, writer io.Writer, addybook *AddressBook) {
	fmt.Fprintln(writer, "1) First name")
	fmt.Fprintln(writer, "2) Last name")
	fmt.Fprintln(writer, "3) Phone number")
	fmt.Fprintln(writer, "4) Address")
	fmt.Fprintln(writer, "5) Email")
	fmt.Fprint(writer, "Field to search: ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	fmt.Fprint(writer, "Value to search: ")
	scanner.Scan()
	value := scanner.Text()
	fmt.Fprintln(writer)
	var addy *Address
	switch choice {
	case 1: addy = addybook.Search("firstname", value)
	case 2: addy = addybook.Search("lastname", value)
	case 3: addy = addybook.Search("phone", value)
	case 4: addy = addybook.Search("address", value)
	case 5: addy = addybook.Search("email", value)
	}
	fmt.Fprintf(writer, "Result: %v\n", addy)
}

func promptUser(scanner *bufio.Scanner, writer io.Writer) int {
	fmt.Fprintln(writer, "1) Load from file")
	fmt.Fprintln(writer, "2) Save to file")
	fmt.Fprintln(writer, "3) Add an entry")
	fmt.Fprintln(writer, "4) Remove an entry")
	fmt.Fprintln(writer, "5) Edit an existing entry")
	fmt.Fprintln(writer, "6) Sort the address book")
	fmt.Fprintln(writer, "7) Search for a specific entry")
	fmt.Fprintln(writer, "8) Quit")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "Please choose what you'd like to do with the database: ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	return choice
}

func manipulateAddressBook(reader io.Reader, writer io.Writer, addybook *AddressBook) {
	scanner := bufio.NewScanner(reader)
	var choice int
	for ok := true; ok; ok = choice != 8 {
		switch choice = promptUser(scanner, writer); choice {
		case 1: loadAddressBook(scanner, writer, addybook)
		case 2: saveAddressBook(scanner, writer, addybook)
		case 3: addEntry(scanner, writer, addybook)
		case 4: removeEntry(scanner, writer, addybook)
		case 5: editEntry(scanner, writer, addybook)
		case 6: sort(scanner, writer, addybook)
		case 7: search(scanner, writer, addybook)
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	addybook := NewAddressBook(NewFileOpener())
	manipulateAddressBook(os.Stdin, os.Stdout, addybook)
}
