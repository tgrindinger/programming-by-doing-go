package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func addKeychains(scanner *bufio.Scanner, writer io.Writer, keychains int) int {
	fmt.Fprintf(writer, "You have %d keychains. How many to add? ", keychains)
	scanner.Scan()
	added, _ := strconv.Atoi(scanner.Text())
	keychains += added
	fmt.Fprintf(writer, "You now have %d keychains.\n", keychains)
	return keychains
}

func removeKeychains(scanner *bufio.Scanner, writer io.Writer, keychains int) int {
	fmt.Fprintf(writer, "You have %d keychains. How many to remove? ", keychains)
	scanner.Scan()
	removed, _ := strconv.Atoi(scanner.Text())
	keychains -= removed
	fmt.Fprintf(writer, "You now have %d keychains.\n", keychains)
	return keychains
}

func viewOrder(writer io.Writer, keychains, price int) {
	fmt.Fprintf(writer, "You have %d keychains.\n", keychains)
	fmt.Fprintf(writer, "Keychains cost $%d each.\n", price)
	fmt.Fprintf(writer, "Total cost is $%d.\n", keychains * price)
}

func checkout(scanner *bufio.Scanner, writer io.Writer, keychains, price int) {
	fmt.Fprintln(writer, "CHECKOUT")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "What is your name? ")
	scanner.Scan()
	name := scanner.Text()
	viewOrder(writer, keychains, price)
	fmt.Fprintf(writer, "Thanks for your order, %s!\n", name)
}

func SellKeychains(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "Ye Olde Keychain Shoppe")
	keychains := 0
	price := 10
	var choice int
	for ok := true; ok; ok = choice != 4 {
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "1. Add Keychains to Order")
		fmt.Fprintln(writer, "2. Remove Keychains from Order")
		fmt.Fprintln(writer, "3. View Current Order")
		fmt.Fprintln(writer, "4. Checkout")
		fmt.Fprintln(writer)
		fmt.Fprint(writer, "Please enter your choice: ")
		scanner.Scan()
		fmt.Fprintln(writer)
		choice, _ = strconv.Atoi(scanner.Text())
		if choice == 1 {
			keychains = addKeychains(scanner, writer, keychains)
		} else if choice == 2 {
			keychains = removeKeychains(scanner, writer, keychains)
		} else if choice == 3 {
			viewOrder(writer, keychains, price)
		} else if choice == 4 {
			checkout(scanner, writer, keychains, price)
		}
	}
}

func main() {
	SellKeychains(os.Stdin, os.Stdout)
}
