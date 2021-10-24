package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Address struct {
	street string
	city   string
	state  string
	zip    int
}

func (a *Address) toString() string {
	return fmt.Sprintf("%s, %s, %s %d", a.street, a.city, a.state, a.zip)
}

func ReadAddress(scanner *bufio.Scanner) *Address {
	addy := &Address{}
	scanner.Scan()
	addy.street = scanner.Text()
	scanner.Scan()
	addy.city = scanner.Text()
	scanner.Scan()
	stateZip := scanner.Text()
	fmt.Sscan(stateZip, &addy.state, &addy.zip)
	return addy
}

func PrintAddresses(writer io.Writer) {
	reader, _ := os.Open("fake-addresses.txt")
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	addybook := make([]*Address, 6)
	for i := 0; i < len(addybook); i++ {
		addybook[i] = ReadAddress(scanner)
	}
	for i := 0; i < len(addybook); i++ {
		fmt.Fprintf(writer, "%d. %s\n", i + 1, addybook[i].toString())
	}
}

func main() {
	PrintAddresses(os.Stdout)
}
