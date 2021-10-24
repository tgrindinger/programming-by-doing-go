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
	uno := ReadAddress(scanner)
	dos := ReadAddress(scanner)
	tre := ReadAddress(scanner)
	cua := ReadAddress(scanner)
	cin := ReadAddress(scanner)
	sei := ReadAddress(scanner)
	fmt.Fprintf(writer, "%s, %s, %s %d\n", uno.street, uno.city, uno.state, uno.zip)
	fmt.Fprintf(writer, "%s, %s, %s %d\n", dos.street, dos.city, dos.state, dos.zip)
	fmt.Fprintf(writer, "%s, %s, %s %d\n", tre.street, tre.city, tre.state, tre.zip)
	fmt.Fprintf(writer, "%s, %s, %s %d\n", cua.street, cua.city, cua.state, cua.zip)
	fmt.Fprintf(writer, "%s, %s, %s %d\n", cin.street, cin.city, cin.state, cin.zip)
	fmt.Fprintf(writer, "%s, %s, %s %d\n", sei.street, sei.city, sei.state, sei.zip)
}

func main() {
	PrintAddresses(os.Stdout)
}
